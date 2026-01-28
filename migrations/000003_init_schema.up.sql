CREATE TABLE public.events (
                               stream_id text,
                               stream_name text,
                               stream_version int,
                               event_id uuid,
                               event_name text,
                               event_data jsonb,
                               occurred_at timestamptz
);

CREATE TABLE public.posts (
                              id uuid PRIMARY KEY,
                              title text,
                              description text,
                              content text,
                              status int,
                              created_at timestamptz,
                              updated_at timestamptz,
                              user_id text
);

CREATE TABLE public.users (
                              id uuid PRIMARY KEY,
                              user_name text,
                              email text,
                              password text,
                              status int,
                              refresh_token_web TEXT,
                              refresh_token_web_at DATETIME,
                              refresh_token_mobile TEXT,
                              refresh_token_mobile_at DATETIME,
                              created_at timestamptz,
                              updated_at timestamptz
);