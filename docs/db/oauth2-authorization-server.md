# OAuth2 Authorization Server

## Context

The platform currently supports two authentication methods:

1. **Basic Auth** — `POST /api/auth/signin` with `base64(login:password)` → JWT
2. **Gitea Token Exchange** — `GET /api/auth/token?token={gitea_access_token}` → JWT

The `@01-edu/api` package (`01-edu/01-edu-api`) uses the second method: a Gitea access token (SHA1) is exchanged for a Hasura JWT, which is then auto-refreshed via `GET /api/auth/refresh` with the `x-jwt-token` header.

While the Gitea token approach avoids repeated password entry, it still has limitations for third-party integrations:

- **No scoped access** — a Gitea token grants full `user` or `admin_read_only` role access, there is no way to limit an app to read-only profile data
- **No user consent flow** — users must manually create tokens in Gitea settings and copy/paste them into apps
- **No standardized discovery** — third-party developers must reverse-engineer the auth flow from `@01-edu/api` source code
- **No revocation per-app** — revoking a Gitea token breaks all apps using it

## Current Architecture

```
┌──────────────┐    Basic Auth     ┌──────────────┐
│   Platform    │ ───────────────► │  /api/auth/   │
│   Login UI   │                   │   signin      │──► JWT
└──────────────┘                   └──────────────┘
                                          ▲
┌──────────────┐  Gitea Token      ┌──────────────┐
│  @01-edu/api │ ───────────────► │  /api/auth/   │
│   package    │                   │   token       │──► JWT
└──────────────┘                   └──────────────┘
                                          │
                                          ▼
                                   ┌──────────────┐
                                   │  /api/auth/   │
                                   │   refresh     │──► new JWT
                                   └──────────────┘
                                          │
                                          ▼
                                   ┌──────────────┐
                                   │   Hasura      │
                                   │   GraphQL     │
                                   └──────────────┘
```

### JWT Payload Structure (existing)

```json
{
  "sub": "674",
  "iat": 1628602069,
  "ip": "86.75.230.26, 172.23.0.2",
  "exp": 1628774869,
  "https://hasura.io/jwt/claims": {
    "x-hasura-allowed-roles": ["user", "admin_read_only"],
    "x-hasura-campuses": "{}",
    "x-hasura-default-role": "admin_read_only",
    "x-hasura-user-id": "674",
    "x-hasura-token-id": "d2283562-5ae5-4f9e-acf9-c1719b8b44b2"
  }
}
```

Hasura only checks the JWT signature and the `https://hasura.io/jwt/claims` namespace. It does not care how the JWT was issued.

### Gitea Token Creation (current method)

```
POST https://git.{domain}/api/v1/users/{username}/tokens
Authorization: Basic base64(username:password)
Content-Type: application/json
Body: { "name": "my-token" }
→ { "sha1": "a6b38ae9dc69c141ccd3b04bbb2d3091f2e2bbb4" }
```

This requires the user's password once, then the SHA1 token is persistent and can be used with `@01-edu/api`.

## Proposed Enhancement

Build an **OAuth2 Authorization Code flow** on top of the existing token exchange infrastructure, so third-party apps never handle passwords or raw Gitea tokens.

### Proposed Flow

```
1. App redirects user to:
   GET /api/oauth/authorize?client_id=X&redirect_uri=Y&response_type=code&scope=profile&state=Z

2. User is already logged in on the platform (or logs in via existing signin)
   → Sees a consent screen: "App X wants to access your profile and XP data"
   → User clicks Authorize

3. Platform redirects back to the app with an authorization code:
   GET {redirect_uri}?code=ABC&state=Z

4. App exchanges code for a scoped JWT (server-side):
   POST /api/oauth/token
   { grant_type: "authorization_code", code: "ABC", client_id: "X", client_secret: "S" }
   → { access_token: "<scoped-jwt>", token_type: "bearer", expires_in: 172800 }

5. App uses the scoped JWT with GraphQL API (same as today)
```

### How It Builds on Existing Infrastructure

The platform already has all the building blocks:

| Existing Component | How OAuth2 Uses It |
|---|---|
| `/api/auth/signin` (Basic Auth → JWT) | User authenticates on the consent screen via existing login |
| `/api/auth/token` (Gitea token → JWT) | OAuth2 token endpoint issues JWTs using the same signing key |
| `/api/auth/refresh` (JWT → new JWT) | OAuth2 refresh tokens use the same refresh mechanism |
| Hasura JWT validation | Scoped JWTs work unchanged — same signature, same claims namespace |
| `@01-edu/api` package | Continues to work as-is for server-to-server integrations |

### New Endpoints

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/api/oauth/authorize` | GET | Authorization + consent screen |
| `/api/oauth/token` | POST | Code exchange, refresh |
| `/api/oauth/revoke` | POST | Token revocation |
| `/api/oauth/userinfo` | GET | OIDC UserInfo |
| `/.well-known/openid-configuration` | GET | OIDC Discovery |

### New Database Tables

```sql
-- Registered OAuth2 clients (managed by campus admins)
CREATE TABLE oauth_client (
    id              SERIAL PRIMARY KEY,
    client_id       VARCHAR(64) UNIQUE NOT NULL,
    client_secret   VARCHAR(256) NOT NULL,  -- bcrypt hashed
    client_name     VARCHAR(128) NOT NULL,
    redirect_uris   TEXT[] NOT NULL,
    allowed_scopes  TEXT[] NOT NULL DEFAULT '{openid,profile}',
    grant_types     TEXT[] NOT NULL DEFAULT '{authorization_code}',
    created_by      INTEGER REFERENCES "user"(id),
    created_at      TIMESTAMP DEFAULT NOW()
);

-- Temporary authorization codes (single-use, 10 min expiry)
CREATE TABLE oauth_authorization_code (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(128) UNIQUE NOT NULL,
    client_id       VARCHAR(64) REFERENCES oauth_client(client_id),
    user_id         INTEGER REFERENCES "user"(id),
    redirect_uri    TEXT NOT NULL,
    scope           TEXT NOT NULL,
    code_challenge  VARCHAR(128),       -- PKCE (RFC 7636)
    code_challenge_method VARCHAR(10),
    expires_at      TIMESTAMP NOT NULL,
    used            BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMP DEFAULT NOW()
);

-- User consent records
CREATE TABLE oauth_consent (
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER REFERENCES "user"(id),
    client_id       VARCHAR(64) REFERENCES oauth_client(client_id),
    scopes_granted  TEXT[] NOT NULL,
    granted_at      TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, client_id)
);
```

### Scope-to-Hasura-Role Mapping

OAuth2 scopes control what data each app can access, mapped to Hasura roles:

| OAuth2 Scope | Data Access | Hasura Role |
|---|---|---|
| `openid` | User ID only | minimal |
| `profile` | login, firstName, lastName, campus, auditRatio | `anonymous` column set |
| `progress:read` | Own progress, results, grades | `user` (filtered) |
| `xp:read` | Own XP transactions | `user` (filtered) |
| `audit:read` | Own audit data | `user` (filtered) |

### JWT Claims (scoped)

Scoped JWTs issued via OAuth2 include both standard OIDC claims and existing Hasura claims:

```json
{
  "iss": "https://((DOMAIN))",
  "sub": "12345",
  "aud": "client_id_here",
  "exp": 1234567890,
  "scope": "openid profile xp:read",
  "https://hasura.io/jwt/claims": {
    "x-hasura-default-role": "user",
    "x-hasura-allowed-roles": ["user"],
    "x-hasura-user-id": "12345",
    "x-hasura-campuses": "{normandie}",
    "x-hasura-token-id": "oauth-xxxxx"
  }
}
```

Same signing key, same Hasura validation — zero configuration change needed on the Hasura side.

## Backward Compatibility

This proposal is **fully additive** with zero breaking changes:

1. **`/api/auth/signin`** — unchanged, continues working for platform UI
2. **`/api/auth/token?token=`** — unchanged, continues working for `@01-edu/api` and Gitea integrations
3. **`/api/auth/refresh`** — unchanged
4. **Hasura `HASURA_GRAPHQL_JWT_SECRET`** — unchanged (same signing key)
5. **`@01-edu/api` package** — continues working as-is
6. **New `/api/oauth/*` endpoints** — additive, no route conflicts

## Precedent in the Codebase

- The `discordToken` table already stores `accessToken` / `refreshToken` / `expiresAt` — OAuth token storage has existed before
- The `user` table has `githubId` / `githubLogin` columns for external identity linking
- The `@01-edu/api` package already implements a token exchange pattern (`Gitea token → JWT`), which is analogous to an OAuth2 token endpoint
- The `hubspot` integration repo (`01-edu/hubspot`) uses `@01-edu/api` for server-to-server auth

## Use Cases

- **Discord bots** — verify student identity and display XP/progress without collecting passwords or Gitea tokens
- **Student portfolio apps** — personal dashboards pulling real-time data
- **Campus tools** — admin dashboards, attendance trackers, analytics
- **Mobile apps** — native apps with proper OAuth2 PKCE flow
- **Third-party integrations** — standardized auth instead of manual Gitea token exchange

## Security Considerations

- **PKCE (RFC 7636)** required for public clients (SPAs, mobile apps)
- **Authorization codes** expire in 10 minutes, single-use
- **Refresh token rotation**
- **Consent screen** — users explicitly approve each app and its requested scopes
- **Per-app revocation** — users can revoke any app from their profile
- **Rate limiting** on token endpoints

## References

- [RFC 6749 — OAuth 2.0](https://datatracker.ietf.org/doc/html/rfc6749)
- [RFC 7636 — PKCE](https://datatracker.ietf.org/doc/html/rfc7636)
- [OpenID Connect Core](https://openid.net/specs/openid-connect-core-1_0.html)
- [Hasura JWT Auth](https://hasura.io/docs/latest/auth/authentication/jwt/)
- [`@01-edu/api` source](https://github.com/01-edu/01-edu-api)
- [`01-edu/hubspot` integration](https://github.com/01-edu/hubspot)
- [Existing auth documentation](./db-authorization.md)
- [Database structure](./database-structure.md)
