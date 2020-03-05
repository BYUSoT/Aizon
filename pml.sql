create database if not exists pml;

use pml;

CREATE TABLE if not exists `users` (
  `id` int PRIMARY KEY,
  `passhash` varchar(128),
  `superuser` tinyint(1),
  `username` varchar(255),
  `fname` varchar(255),
  `lname` varchar(255),
  `email` varchar(255),
  `staff` tinyint(1)
);

CREATE TABLE if not exists `opunits` (
  `id` int PRIMARY KEY,
  `unit` varchar(255),
  `account` varchar(255),
  `class` varchar(255),
  `cust_depid` int
);

CREATE TABLE if not exists `payments` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE if not exists `orders` (
  `id` int PRIMARY KEY,
  `cust_depid` int,
  `ordered` date,
  `completed` date,
  `cpo` varchar(255),
  `team` varchar(255),
  `desc` text,
  `state` int,
  `auxcontactid` int,
  `billcontactid` int,
  `opunitid` int,
  `pmtid` int
);

CREATE TABLE if not exists `states` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE if not exists `contacts` (
  `id` int PRIMARY KEY,
  `fname` varchar(255),
  `lname` varchar(255),
  `email` varchar(255),
  `address1` varchar(255),
  `address2` varchar(255),
  `city` varchar(255),
  `state` varchar(255),
  `zip` varchar(255)
);

CREATE TABLE if not exists `departments` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `orgid` int,
  `rateid` int
);

CREATE TABLE if not exists `customers` (
  `id` int PRIMARY KEY,
  `contactid` int,
  `foreign_id` varchar(255)
);

CREATE TABLE if not exists `customer_dept` (
  `id` int PRIMARY KEY,
  `custid` int,
  `depid` int
);

CREATE TABLE if not exists `orgs` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE if not exists `rates` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `cents` int,
  `islabor` tinyint(1),
  `staff` tinyint(1)
);

CREATE TABLE if not exists `hours` (
  `id` int PRIMARY KEY,
  `orderid` int,
  `userid` int,
  `rateid` int,
  `qrtrhour` int,
  `desc` text,
  `when` date
);

CREATE TABLE if not exists `materials` (
  `id` int PRIMARY KEY,
  `orderid` int,
  `userid` int,
  `itemid` int,
  `quantity` int,
  `desc` text,
  `dimension` text,
  `price` int,
  `when` date
);

CREATE TABLE if not exists `items` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);
