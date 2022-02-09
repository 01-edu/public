## Introduction

Following are examples of SQL queries to help you get you started.

## Examples of SQL queries

### Sign up

The following query returns the users that are currently in the sign up process.

```sql
SELECT DISTINCT
    u.id,
    u."githubLogin",
    u.attrs ->> 'email' as email,
    u.attrs ->> 'image' as image,
    u.attrs ->> 'lastName' as "last-name",
    u.attrs ->> 'firstName' as "firstName",
    u."createdAt" as "account-creation-date"
FROM public.user u
WHERE u.id NOT IN (SELECT "userId" FROM public.user_role)
AND u.id NOT IN (SELECT "userId" FROM public.progress p WHERE p."userId"=u.id);
```

### Toad

The following query returns the users that are currently in the toad process. The information includes user and games information.

```sql
-- query for **TOAD**
SELECT DISTINCT
    u.id,
    "githubLogin",
    u.attrs ->> 'name' as name,
    u.attrs ->> 'email' as email,
    u.attrs ->> 'image' as image,
    u.attrs ->> 'gender' as gender,
    u.attrs ->> 'country' as country,
    u.attrs ->> 'language' as language,
    u.attrs ->> 'lastName' as "last-name",
    u.attrs ->> 'firstName' as "firstName",
    u.attrs ->> 'dateOfBirth' as "date-of-birth",
    u.attrs ->> 'environment' as environment,
    u.attrs ->> 'general-conditionsAcceptedname' as "general-conditions-accepted",
    u."createdAt" as "account-creation-date",
    r.attrs ->> 'score' as "game-score",
    r.attrs ->> 'allowedAttempts' as "allowed-attempts",
    r.attrs ->> 'games' as games
FROM public.user u
LEFT JOIN public.result r ON u.id=r."userId" AND r.path LIKE '%/onboarding/games'
WHERE r."type"='admin_selection'
AND u.id NOT IN (SELECT "userId" FROM public.user_role)
AND r.grade IS NULL
AND r.attrs ->> 'games' IS NOT NULL
ORDER BY r.attrs ->> 'score' ASC;
```

The following query returns all the users who signed in to the platform at least once and have never started playing the games.

```sql
SELECT *
FROM public.user
WHERE ((NOT ((EXISTS (
          SELECT 1
          FROM public.progress AS p
          WHERE p."userId" = public.user.id)))
      AND (((NOT (EXISTS (
                SELECT 1
                FROM toad.sessions AS s
                WHERE s."candidate_id" = public.user.id)
                OR (EXISTS (
                    SELECT 1
                    FROM toad.sessions AS s
                    WHERE (s.candidate_id = public.user.id
                      AND s.started_at IS NULL)))))
            AND (NOT ((EXISTS (
                    SELECT 1
                    FROM public.user_roles_view AS r
                    WHERE (r."userId" = public.user.id
                      AND r.slug IN ('admin', 'campus_admin_gritlab'))))))))))
```

### Administration

The following query returns the users currently in the administration process. This includes all information about the user, the number of attempts and the current phase.

```sql
-- query for **administration**
SELECT DISTINCT
    u.id,
    u."githubLogin",
    r.attrs ->> 'name' as name,
    r.attrs ->> 'email' as email,
    r.attrs ->> 'image' as image,
    r.attrs ->> 'gender' as gender,
    r.attrs ->> 'country' as country,
    r.attrs ->> 'language' as language,
    r.attrs ->> 'firstName' as "first-name",
    r.attrs ->> 'lastName' as "last-name",
    r.attrs ->> 'discordId' as "discord-id",
    r.attrs ->> 'addressCity' as "address-city",
    r.attrs ->> 'dateOfBirth' as "date-of-birth",
    r.attrs ->> 'environment' as environment,
    r.attrs ->> 'medicalInfo' as "medical-info" ,
    r.attrs ->> 'discordLogin' as "discord-login" ,
    r.attrs ->> 'placeOfBirth' as "place-of-birth" ,
    r.attrs ->> 'addressStreet' as "address-street" ,
    r.attrs ->> 'addressCountry' as "address-country" ,
    r.attrs ->> 'countryOfBirth' as "country-of-birth" ,
    r.attrs ->> 'chart01Accepted' as "chart-01-accepted" ,
    r.attrs ->> 'id-cardUploadId' as "id-card-uploadId",
    r.attrs ->> 'addressPostalCode' as "address-postal-code" ,
    r.attrs ->> 'emergencyLastName' as "emergency-last-name" ,
    r.attrs ->> 'emergencyFirstName' as "emergency-first-name" ,
    r.attrs ->> 'regulationAccepted' as "regulation-accepted" ,
    r.attrs ->> 'emergencyAffiliation' as "emergency-affiliation" ,
    r.attrs ->> 'addressComplementStreet' as "address-complement-street" ,
    r.attrs ->> 'general-conditionsAccepted' as "general-conditions-accepted",
    u."createdAt" as "account-creation-date",
    r.attrs ->> 'phase' as phase,
    COUNT(p.id) as attempts
FROM public.user u
LEFT JOIN public.result r ON u.id=r."userId" AND r.path LIKE '%/onboarding/administration' AND r.grade IS NULL
LEFT JOIN public.progress p ON u.id=p."userId" AND p.path LIKE '%/onboarding/administration'
WHERE r."type"='admin_selection'
AND u.id NOT IN (SELECT "userId" FROM public.user_role)
GROUP BY u.id, r.attrs;
```

### Xp per user per activity

The following query returns the amount of xp per user and per activity.

```sql
WITH xp_user AS (
    SELECT
        u."githubLogin",
        xp.amount,
        xp.path,
        xp."eventParentId"
    FROM public.user u
    LEFT JOIN public.xp_by_event xp ON xp."userId"=u.id
    WHERE xp.amount IS NOT NULL
    ORDER BY u."githubLogin" ASC
)
SELECT
    xu."githubLogin",
    xu.amount,
    xu.path,
    e.path as "parent-path"
FROM xp_user xu
LEFT JOIN public.event e ON e.id=xu."eventParentId";
```

### Xp per user

The following query returns the amount of xp per user.

```sql
-- user per xp
SELECT
    u.id,
    u."githubLogin",
    u.attrs ->> 'name' as name,
    u.attrs ->> 'email' as email,
    u.attrs ->> 'image' as image,
    u.attrs ->> 'gender' as gender,
    u.attrs ->> 'country' as country,
    u.attrs ->> 'language' as language,
    u.attrs ->> 'lastName' as "last-name",
    u.attrs ->> 'firstName' as "firstName",
    u.attrs ->> 'dateOfBirth' as "date-of-birth",
    u.attrs ->> 'environment' as environment,
    u.attrs ->> 'general-conditionsAcceptedname' as "general-conditions-accepted",
    u."createdAt" as "account-creation-date",
    xp.amount as "xp-amount"
FROM public.user u, public.xp xp
WHERE u.id=xp."userId"
ORDER BY xp.amount DESC;
```

### Group status per captain

The following query returns the groups status per captain.

```sql
-- group status
SELECT
    u."githubLogin",
    g.path,
    g.status
FROM public.user u, public."group" g
WHERE u.id=g."captainId"
ORDER BY u."githubLogin" ASC;
```

### Group progresses

The following query returns the progress per group.

```sql
-- group progresses
WITH progress_group AS (
    SELECT
        p.path,
        p.grade,
        p."isDone",
        p.campus,
        g."captainId",
        g.status
    FROM public.progress p
    LEFT JOIN public."group" g ON g.id=p."groupId"
)
SELECT DISTINCT
    u."githubLogin",
    pg.status as "group-status",
    pg.path,
    pg.grade,
    pg."isDone",
    pg.campus
FROM public.user u
LEFT JOIN progress_group pg ON "captainId"=u.id
WHERE pg.path IS NOT NULL
ORDER BY u."githubLogin" ASC;
```
