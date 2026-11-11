create database crowdfunding_sandbox

create table donors (
	id serial primary key,
	name varchar(100) not null,
	email varchar(255) not null unique,
	password varchar(255) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp default current_timestamp
)

create table admins (
	id serial primary key,
	name varchar(100) not null,
	email varchar(255) not null unique,
	password varchar(255) not null,
	role varchar(20),
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp default current_timestamp
)

create table foundations (
	id serial primary key,
	name varchar(100) not null,
	slug varchar(255) unique,
	desription text,
	address varchar(255) not null,
	logo_url varchar(500),
	admin_id serial not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp default current_timestamp,
	foreign key(admin_id) references admins(id)
)

create type campaign_status as enum ('draft', 'active', 'completed', 'cancelled'); 

create table campaigns (
	id serial primary key,
	foundation_id serial not null,
	title varchar(255) not null,
	slug varchar(255),
	desription text,
	target_amount NUMERIC(15, 2) NOT null,
  	current_amount NUMERIC(15, 2) NOT null,
  	status campaign_status,
  	start_date TIMESTAMP NOT null,
  	end_date TIMESTAMP NOT null,
 	created_by serial NOT null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp default current_timestamp,
	foreign key(created_by) references admins(id)
);

comment on column campaigns.status is 'draft, active, completed, cancelled'

alter table campaigns
add constraint FK_foundation_campaign
foreign key (foundation_id) references foundations(id) on delete cascade on update cascade

create type donation_status as enum ('pending', 'success', 'failed', 'expired'); 

create table donations (
    id serial PRIMARY key,
    donor_id serial NOT null,
    campaign_id serial NOT null,
    invoice_number VARCHAR(100) NOT NULL unique,
    amount NUMERIC(15, 2) NOT null,
    status donation_status,
    message TEXT,
    created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp default current_timestamp,
	constraint FK_donor foreign key (donor_id) references donors(id) 
	on delete cascade on update cascade,
	constraint FK_campaingn foreign key (campaign_id) references campaigns(id) 
	on delete cascade on update cascade 
)

comment on column donations.status is 'pending, success, completed, cancelled'

create TABLE payment_logs (
    id serial PRIMARY key,
    donation_id serial NOT null,
    gateway_name VARCHAR(50) NOT null, -- Misal: Midtrans, Xendit, Stripe
    payment_method VARCHAR(50), -- Misal: qris, bca_va, gopay
    paid_at TIMESTAMP,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP default current_timestamp,
    constraint fk_donation foreign key (donation_id) references donations(id)
    on delete cascade on update cascade
)
