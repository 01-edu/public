# OAuth2 Authorization Server

## Context

Currently, the only way for external applications to authenticate a user is via `POST /api/auth/signin` with Basic Auth (`base64(login:password)`). This requires third-party applications to **collect and transmit user passwords directly**, which creates security risks and trust barriers.

An OAuth2 Authorization Server would allow third-party apps (Discord bots, student dashboards, portfolio tools, mobile apps) to authenticate users **without ever seeing their password**.

## Current Limitations

- **Credential exposure**: Third-party apps must handle raw passwords
- **No granular permissions**: Any valid JWT grants full `user` role access
- **No revocation**: Users cannot revoke a specific app's access without changing their password
- **No audit trail**: No distinction between API calls from the user vs. from a third-party app

## Proposed Architecture

### Authentication Flow

```
1. App redirects user to:
   GET /api/oauth/authorize?client_id=X&redirect_uri=Y&response_type=code&scope=openid+profile&state=Z

2. User sees a consent screen on the 01-edu platform
   (already logged in, or logs in via the existing signin form)

3. Platform redirects back to the app with an authorization code:
   GET {redirect_uri}?code=ABC&state=Z

4. App exchanges the code for a token (server-side):
   POST /api/oauth/token
   { grant_type: "authorization_code", code: "ABC", client_id: "X", client_secret: "S" }

5. App receives a scoped JWT

6. App calls the GraphQL API with the scoped JWT
```

The user's password **never leaves the 01-edu platform**.

### New Endpoints

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/api/oauth/authorize` | GET | Authorization + consent screen |
| `/api/oauth/token` | POST | Code exchange, refresh tokens |
| `/api/oauth/revoke` | POST | Token revocation |
| `/api/oauth/userinfo` | GET | OIDC UserInfo |
| `/.well-known/openid-configuration` | GET | OIDC Discovery |
| `/.well-known/jwks.json` | GET | Public key for JWT verification |

### New Database Tables

```sql
-- Registered OAuth2 clients (managed by admins)
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
    code_challenge  VARCHAR(128),       -- PKCE support (RFC 7636)
    code_challenge_method VARCHAR(10),  -- S256
    expires_at      TIMESTAMP NOT NULL,
    used            BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMP DEFAULT NOW()
);

-- User consent records (what apps a user has authorized)
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

OAuth2 scopes control what data each app can access:

| OAuth2 Scope | Data Access | Hasura Equivalent |
|---|---|---|
| `openid` | User ID only | minimal |
| `profile` | login, firstName, lastName, campus, auditRatio | `anonymous` column set |
| `progress` | Own progress, results, grades | `user` filtered |
| `xp` | Own XP transactions | `user` filtered |
| `audit` | Own audit data | `user` filtered |

### JWT Claims

The JWT issued via OAuth2 includes both standard OIDC claims and Hasura claims:

```json
{
  "iss": "https://((DOMAIN))",
  "sub": "12345",
  "aud": "client_id_here",
  "exp": 1234567890,
  "scope": "openid profile xp",
  "https://hasura.io/jwt/claims": {
    "x-hasura-default-role": "user",
    "x-hasura-allowed-roles": ["user"],
    "x-hasura-user-id": "12345",
    "x-hasura-campuses": "{campus-name}"
  }
}
```

Hasura only validates the JWT signature and the `https://hasura.io/jwt/claims` namespace, so this is fully backward-compatible.

## Backward Compatibility

This proposal is **fully additive** with zero breaking changes:

1. **`/api/auth/signin`** continues working as-is for the platform UI
2. **New `/api/oauth/*` endpoints** do not conflict with existing routes
3. **Same JWT format** — Hasura does not care how the JWT was issued
4. **Same database** — new tables alongside existing ones, no migrations on existing tables
5. **`HASURA_GRAPHQL_JWT_SECRET`** configuration remains unchanged

## Precedent

The platform has handled OAuth tokens before:

- The `discordToken` table already stores `accessToken` / `refreshToken` / `expiresAt`
- The `user` table has `githubId` / `githubLogin` columns for GitHub identity linking

## Use Cases

- **Discord bots**: Verify student identity and display XP/progress without collecting passwords
- **Student portfolios**: Personal dashboard apps pulling real-time data from the platform
- **Campus tools**: Admin dashboards, attendance trackers, analytics
- **Mobile apps**: Native apps with OAuth2 PKCE flow
- **CI/CD integrations**: Automated tools acting on behalf of students

## Security Considerations

- **PKCE (RFC 7636)** required for public clients (SPAs, mobile apps)
- **Authorization codes** expire in 10 minutes, single-use
- **Refresh token rotation** — each refresh issues a new token
- **Consent screen** — users explicitly approve each app and its requested scopes
- **Revocation** — users can revoke any app's access from their profile settings
- **Rate limiting** on all token endpoints

## References

- [RFC 6749 — OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749)
- [RFC 7636 — Proof Key for Code Exchange (PKCE)](https://datatracker.ietf.org/doc/html/rfc7636)
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)
- [Hasura JWT Authentication](https://hasura.io/docs/latest/auth/authentication/jwt/)
- [Existing db-authorization.md](./db-authorization.md)
- [Existing database-structure.md](./database-structure.md)
