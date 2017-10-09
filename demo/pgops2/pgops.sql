--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.4
-- Dumped by pg_dump version 9.6.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: test; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE test (
    gopher_id integer,
    created timestamp with time zone
);


--
-- Data for Name: test; Type: TABLE DATA; Schema: public; Owner: -
--

COPY test (gopher_id, created) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

