--
-- PostgreSQL database dump
--

-- Dumped from database version 10.19 (Debian 10.19-1.pgdg90+1)
-- Dumped by pg_dump version 10.19 (Debian 10.19-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = true;

--
-- Name: criminal_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.criminal_records (
    id integer NOT NULL,
    name character varying(120) NOT NULL,
    felony character varying(30) NOT NULL,
    ssn character varying(11) NOT NULL,
    home_address character varying(100) NOT NULL,
    entry timestamp without time zone NOT NULL,
    city character varying(100) NOT NULL,
    status character varying(16) NOT NULL
);


ALTER TABLE public.criminal_records OWNER TO postgres;

--
-- Name: criminal_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.criminal_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.criminal_records_id_seq OWNER TO postgres;

--
-- Name: criminal_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.criminal_records_id_seq OWNED BY public.criminal_records.id;


--
-- Name: criminal_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.criminal_records ALTER COLUMN id SET DEFAULT nextval('public.criminal_records_id_seq'::regclass);


--
-- Data for Name: criminal_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.criminal_records (id, name, felony, ssn, home_address, entry, city, status) FROM stdin;
1	Anna Strickland	Larceny	619-12-6717	70189 Michael Mall Apt. 896	2003-09-01 00:00:00	North Taraton, NJ 22837	missing
2	Thomas Jimenez	Tax evasion	188-95-4383	1991 Wilkins Avenue	2021-06-09 00:00:00	Singhberg, ID 97780-0873	terminated
3	Jamie Mitchell	Obstruction of justice	717-08-1960	219 Harris Mill Suite 207	2015-01-14 00:00:00	Collinsburgh, WY 89313	missing
4	Edward Snyder	Animal cruelty	092-62-0760	8132 Bradley Curve	1987-06-22 00:00:00	Jacksonstad, OR 01263-1524	alive
5	Linda Khan	Animal cruelty	782-95-0048	21299 Joseph Route Suite 773	2008-02-17 00:00:00	Raymondburgh, MN 12498	terminated
6	Danny Rowe	Perjury	216-78-3132	09404 Lutz Extensions	1990-11-30 00:00:00	Kellyhaven, MP 95620-1505	alive
7	Amanda Patterson	Check fraud	230-04-2542	7372 Jason Crest Suite 693	1982-09-23 00:00:00	Matthewfort, MI 14415-5009	missing
8	Jacqueline Hubbard	Arson	460-04-4203	2341 Benjamin Squares Suite 075	2006-09-11 00:00:00	North Jeremymouth, NJ 50524	alive
9	Samuel Hunt	Vehicular homicide	582-43-8485	559 Tyler Overpass Suite 701	1973-05-08 00:00:00	Joelfurt, MH 62837-0667	terminated
10	Erica Moore	Burglary	213-17-3185	27267 Lee Vista	2003-02-26 00:00:00	Contreraston, WV 81915-2211	alive
11	Terri Soto	Vehicular homicide	237-73-2165	380 Bennett Lodge Suite 097	1992-04-21 00:00:00	Williamsberg, RI 26884-5394	missing
12	Mr. Curtis Livingston Jr.	Check fraud	065-22-2093	966 Hanna Dam	1996-12-23 00:00:00	East Tyler, NE 98319	missing
13	Glenda Larsen	Check fraud	358-81-9699	49868 Jesse Meadows Suite 886	2011-11-29 00:00:00	Laurafurt, AK 23242-6157	terminated
14	Kevin Smith	Manslaughter	113-47-7791	39906 Marcus Knolls Suite 239	1999-01-08 00:00:00	South Briannastad, PA 33510	missing
15	James Cowan	Perjury	464-39-9838	Unit 4123 Box 7730	1995-03-07 00:00:00	DPO AP 05196	alive
16	Dr. Tiffany Trujillo	Burglary	739-95-7853	92892 Kelly Forest Apt. 243	1971-11-26 00:00:00	Franklinton, KS 37940	missing
17	Vincent Lynch	Perjury	344-20-8764	20234 Miller Cove Apt. 164	1971-04-22 00:00:00	Brownchester, ME 13069	missing
18	Donald Charles	Tax evasion	021-08-7099	14541 Cynthia Throughway Apt. 863	2006-08-11 00:00:00	Angelicaburgh, AR 37085-0775	missing
19	Susan Bates	Obstruction of justice	348-22-6540	382 Gibbs Underpass Suite 276	1990-08-12 00:00:00	East Robert, RI 78390	alive
20	Douglas Mcguire	Perjury	511-76-0710	0665 Jason Flats	2004-03-27 00:00:00	Clarkstad, NE 49111-7058	alive
21	Jermaine Williams	Larceny	474-81-7203	3985 Elizabeth Islands Suite 909	1984-12-30 00:00:00	West Nancy, TX 44604	missing
22	Wendy Hartman	Check fraud	334-42-9907	390 Smith Mills Suite 909	1985-01-24 00:00:00	Williamfort, AK 72811-4233	terminated
23	James Mcdonald	Manslaughter	811-63-9692	4054 Aaron Forks Apt. 577	2020-05-23 00:00:00	Salinasshire, NC 31101-3577	missing
24	Michael Flynn	Manslaughter	222-87-6760	82487 Casey Pine Apt. 198	2015-06-09 00:00:00	Cooperland, NC 02045	terminated
25	Jacob Mathews	Animal cruelty	661-45-7236	2740 Boyd Rest	1998-05-02 00:00:00	Jamesview, OK 68952	alive
26	Madeline Brady	Burglary	179-95-1431	4718 Matthew Gateway	1999-11-19 00:00:00	Hicksview, UT 77869	missing
27	Ashley Hall	Obstruction of justice	806-93-5490	811 Robert Street Suite 365	1988-09-17 00:00:00	East Gloriaside, NJ 19279	terminated
28	Amy Ellison	Vehicular homicide	391-70-2449	998 Rebekah Crest Suite 437	1970-10-18 00:00:00	Richardsburgh, MT 20115	alive
29	Christina Gilbert	Check fraud	634-04-0207	61168 Scott Burg Apt. 909	1977-05-15 00:00:00	Owensmouth, AL 56654-0303	missing
30	Anthony Walls	Obstruction of justice	718-11-9213	31869 James Spur	1974-01-06 00:00:00	Lake Melissa, AS 76349-9034	terminated
31	Hannah Holland	Vehicular homicide	070-08-4524	USNS Pitts	2008-08-31 00:00:00	FPO AA 73968	terminated
32	Joseph Perez	Obstruction of justice	292-45-7374	36369 Graham Ferry Apt. 825	2005-07-17 00:00:00	Bakerstad, WY 73251-7312	alive
33	Lisa Bryant	Obstruction of justice	618-99-5148	315 Sarah Plaza	2004-12-04 00:00:00	Lake Tina, KS 62859	missing
34	Jennifer Arellano	Obstruction of justice	011-03-8345	2011 Steven Tunnel Suite 107	2013-03-08 00:00:00	Port Susanmouth, AR 78234-6943	terminated
35	Brent Davis	Arson	841-82-3085	Unit 4115 Box 6874	1983-08-06 00:00:00	DPO AE 02631	terminated
36	Chad Rodgers	Arson	340-80-8330	3979 William Lodge Suite 097	2010-07-06 00:00:00	Victoriaborough, AZ 35759-0102	missing
37	Karina Grant	Check fraud	882-58-6785	1027 Brandon Tunnel	2004-12-30 00:00:00	Lake Jaredmouth, MT 05088	terminated
38	Mary Sanchez	Arson	277-05-3164	25700 Danielle Meadows	1970-02-09 00:00:00	Ingramview, OH 33911	alive
39	Thomas Hester	Vehicular homicide	607-09-2820	2329 Torres Expressway	1997-09-21 00:00:00	South Yolanda, DE 16847	missing
40	Jessica Moore	Check fraud	430-89-1267	Unit 8067 Box 3284	1997-01-13 00:00:00	DPO AE 14473	alive
41	Kayla Cobb	Obstruction of justice	873-18-1418	583 Kathy Path Apt. 908	2015-03-01 00:00:00	Cookside, OH 00082	terminated
42	Steven Archer	Check fraud	731-32-0542	53715 Sherri Bypass	1986-01-02 00:00:00	Larryfort, FL 08738	alive
43	Brendan Chavez	Larceny	330-94-0203	475 Alejandro Trafficway	2014-10-28 00:00:00	South Daniel, IN 17497-1471	alive
44	Antonio Raymond	Check fraud	408-22-6881	25945 Wheeler Lock	1988-01-13 00:00:00	Paulshire, VT 86439	terminated
45	Katherine Bowman	Burglary	671-68-1186	263 Hernandez Ville Apt. 377	1997-03-09 00:00:00	West Susanville, WI 73920-8166	alive
46	Sarah Terry	Perjury	536-03-1815	USNS Porter	1984-09-09 00:00:00	FPO AP 74010-9771	missing
47	Kevin Nguyen	Perjury	781-56-7455	38399 Kelly Cape Apt. 426	2013-05-24 00:00:00	Hahnstad, IL 41989-1162	alive
48	Zoe Wilson	Vehicular homicide	017-54-4916	1173 James Haven Suite 789	1980-11-30 00:00:00	Williamston, VT 48765-7701	alive
49	Sarah Hart	Animal cruelty	594-01-0712	39834 Megan Views	1995-03-25 00:00:00	South Shaneport, KY 12282-0355	terminated
50	Sara Walters	Larceny	883-07-1450	USCGC Mills	2001-12-04 00:00:00	FPO AP 14523	missing
51	James Ball	Arson	307-74-7238	85520 Johnson Fords	1989-07-07 00:00:00	West Stephaniemouth, TN 11317-8392	alive
52	Andrea Mckinney	Animal cruelty	301-91-3323	8895 Jared Gardens Apt. 939	2001-11-24 00:00:00	Jamesville, WA 69236-9075	alive
53	Melissa Simmons	Arson	880-28-6055	042 Floyd Villages Suite 430	2003-10-12 00:00:00	Lake Amy, MP 05293-3198	missing
54	Willie Wilson	Tax evasion	507-55-0717	6582 Anderson Mountain	1979-11-06 00:00:00	Lauraland, AL 88451	missing
55	Lisa Bryant	Perjury	359-64-8371	Unit 3975 Box 0587	2007-09-16 00:00:00	DPO AE 05621	alive
56	Justin Johnson	Larceny	436-76-2501	595 Henry Trace Suite 552	2016-02-27 00:00:00	North Josephland, MD 04953-8044	missing
57	Rebecca Beck	Larceny	331-95-3692	89953 Castaneda Drives	1990-03-15 00:00:00	North Shane, AR 74035-7871	alive
58	Ann Brown	Arson	527-51-6412	590 Sullivan Lodge	1986-06-11 00:00:00	East Nathanborough, PR 01181-5325	missing
59	Nancy Harper	Animal cruelty	695-96-7190	892 Gregory Tunnel Apt. 069	1994-11-09 00:00:00	Rhondabury, VT 90171	missing
60	Lisa Snyder	Animal cruelty	006-13-3052	38209 Teresa Village	2010-12-28 00:00:00	Dannyhaven, AK 10235-1280	alive
61	Kathleen Rubio	Tax evasion	131-53-4775	15380 Patricia Light Apt. 518	2017-12-10 00:00:00	Michaelbury, PR 27405-1015	terminated
62	Eric Simon	Check fraud	321-32-5658	11542 Simpson Station Suite 625	2019-07-21 00:00:00	Stephanieburgh, NE 80853-0231	missing
63	Janice Meyer	Check fraud	226-85-4073	30214 Kristin Crest Suite 897	1995-11-15 00:00:00	Carolmouth, NM 46656-4754	alive
64	Randy Barry	Tax evasion	414-30-1227	2140 Seth Mount Apt. 424	1982-03-19 00:00:00	Brownside, CO 86683	terminated
65	Bradley Long	Obstruction of justice	733-54-0998	PSC 3442, Box 7315	1971-04-18 00:00:00	APO AE 90593-4537	terminated
66	Connie Keller	Obstruction of justice	026-65-8039	473 Cook Ridges	2002-07-18 00:00:00	New Erichaven, IN 08777	alive
67	Joseph Perez	Arson	120-09-7475	USCGC Allen	1984-10-17 00:00:00	FPO AA 78003	missing
68	Gary Owens	Arson	684-33-7190	62777 Carroll Center Suite 672	1997-12-25 00:00:00	West Davidville, CT 87822-4729	alive
\.


--
-- Name: criminal_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.criminal_records_id_seq', 68, true);


--
-- Name: criminal_records criminal_records_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.criminal_records
    ADD CONSTRAINT criminal_records_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

