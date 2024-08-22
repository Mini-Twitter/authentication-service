create table users(
    id uuid default gen_random_uuid() primary key,
    first_name varchar,
    last_name varchar,
    email varchar unique,
    phone varchar(13)unique ,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint default 0,
    unique(email, deleted_at)
);

create table user_profile(
    user_id uuid references users(id),
    username varchar unique,
    password varchar,
    nationality varchar,
    bio varchar,
    profile_image varchar,
    followers_count int default 0,
    following_count int default 0,
    posts_count int default 0,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    is_active bool
);

create table follows(
    follower_id uuid references users(id),
    following_id uuid references users(id),
    created_at timestamp default now()
)