CREATE TABLE public.events (
                               stream_id text,
                               stream_name text,
                               stream_version text,
                               event_id text,
                               event_name text,
                               event_data text,
                               occurred_at timestamp with time zone
);

CREATE TABLE public.posts (
                              id text PRIMARY KEY,
                              title text,
                              description text,
                              content text,
                              status bigint,
                              created_at timestamp with time zone,
                              updated_at timestamp with time zone,
                              user_id text
);