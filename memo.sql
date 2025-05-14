/* public.users */
create table public.users (
  id uuid primary key references auth.users on delete cascade,
  github_name text,
  avatar_url text,
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);

create or replace function public.handle_new_user()
returns trigger as $$
begin
  insert into public.users (uid, github_name, avatar_url)
  values (
    new.id,
    new.raw_user_meta_data ->> 'user_name',
    new.raw_user_meta_data ->> 'avatar_url'
  );
  return new;
end;
$$ language plpgsql security definer;

create trigger on_auth_user_created
after insert on auth.users
for each row
execute function public.handle_new_user();

/* public.rooms */
create table public.rooms (
  id uuid primary key default gen_random_uuid(),
  host_id uuid references public.users(id) on delete cascade,
  name text not null,
  capacity int not null,
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);