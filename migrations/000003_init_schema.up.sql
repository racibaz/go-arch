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
                              created_at timestamptz,
                              updated_at timestamptz
);

CREATE TABLE public.refresh_tokens (
                                       user_id uuid REFERENCES public.users(id) ON DELETE CASCADE,
                                       hashed_token varchar(500) NOT NULL UNIQUE,
                                       created_at timestamptz,
                                       expires_at timestamptz,
                                       PRIMARY KEY (user_id, hashed_token)
);
