create table users(
    id uuid default gen_random_uuid() primary key,
    first_name varchar,
    last_name varchar,
    email varchar unique,
    email_password varchar,
    phone varchar(13),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0
);

create table user_profile(
    user_id uuid references users(id),
    username varchar unique,
    user_password varchar,
    nationality varchar,
    bio varchar,
    profile_image varchar,
    followers_count int,
    following_count int,
    posts_count int,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    is_active bool
);

create table follows(
    follower_id uuid references users(id),
    following_id uuid references users(id),
    created_at timestamp default now()
)