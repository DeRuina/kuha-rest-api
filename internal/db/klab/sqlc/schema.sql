--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0 (Debian 17.0-1.pgdg110+1)
-- Dumped by pg_dump version 17.5 (Ubuntu 17.5-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: customer; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.customer (
    idcustomer integer NOT NULL,
    firstname character varying(45) NOT NULL,
    lastname character varying(45) NOT NULL,
    idgroups integer,
    dob date,
    sex integer,
    dob_year integer,
    dob_month integer,
    dob_day integer,
    pid_number character varying(45),
    company character varying(45),
    occupation character varying(45),
    education character varying(45),
    address character varying(145),
    phone_home character varying(45),
    phone_work character varying(45),
    phone_mobile character varying(45),
    faxno character varying(45),
    email character varying(150),
    username character varying(45),
    password character varying(45),
    readonly integer,
    warnings integer,
    allow_to_save integer,
    allow_to_cloud integer,
    flag2 integer,
    idsport integer,
    medication text,
    addinfo text,
    team_name character varying(64),
    add1 integer,
    athlete integer,
    add10 character varying(45),
    add20 character varying(45),
    updatemode integer,
    weight_kg real,
    height_cm real,
    date_modified double precision,
    recom_testlevel integer,
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint,
    allow_anonymous_data character varying(45),
    locked smallint,
    allow_to_sprintai integer,
    tosprintai_from date,
    stat_sent date,
    sportti_id character varying(64)
);


ALTER TABLE public.customer OWNER TO klabadmin;

--
-- Name: customer_idcustomer_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.customer_idcustomer_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customer_idcustomer_seq OWNER TO klabadmin;

--
-- Name: customer_idcustomer_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.customer_idcustomer_seq OWNED BY public.customer.idcustomer;


--
-- Name: dirrawdata; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.dirrawdata (
    iddirrawdata integer NOT NULL,
    idmeasurement integer NOT NULL,
    rawdata text,
    columndata character varying(100),
    info character varying(100),
    unitsdata character varying(100),
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint
);


ALTER TABLE public.dirrawdata OWNER TO klabadmin;

--
-- Name: dirrawdata_iddirrawdata_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.dirrawdata_iddirrawdata_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dirrawdata_iddirrawdata_seq OWNER TO klabadmin;

--
-- Name: dirrawdata_iddirrawdata_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.dirrawdata_iddirrawdata_seq OWNED BY public.dirrawdata.iddirrawdata;


--
-- Name: dirreport; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.dirreport (
    iddirreport integer NOT NULL,
    page_instructions text,
    idmeasurement integer NOT NULL,
    template_rec integer,
    librec_name character varying(50),
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint
);


ALTER TABLE public.dirreport OWNER TO klabadmin;

--
-- Name: dirreport_iddirreport_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.dirreport_iddirreport_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dirreport_iddirreport_seq OWNER TO klabadmin;

--
-- Name: dirreport_iddirreport_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.dirreport_iddirreport_seq OWNED BY public.dirreport.iddirreport;


--
-- Name: dirresults; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.dirresults (
    iddirresults integer NOT NULL,
    idmeasurement integer NOT NULL,
    max_vo2mlkgmin double precision,
    max_vo2mlmin double precision,
    max_vo2 double precision,
    max_hr double precision,
    max_speed double precision,
    max_pace double precision,
    max_p double precision,
    max_pkg double precision,
    max_angle double precision,
    max_lac double precision,
    max_add1 double precision,
    max_add2 double precision,
    max_add3 double precision,
    lac_ank_vo2mlkgmin double precision,
    lac_ank_vo2mlmin double precision,
    lac_ank_vo2 double precision,
    lac_ank_vo2pr double precision,
    lac_ank_hr double precision,
    lac_ank_speed double precision,
    lac_ank_pace double precision,
    lac_ank_p double precision,
    lac_ank_pkg double precision,
    lac_ank_angle double precision,
    lac_ank_lac double precision,
    lac_ank_add1 double precision,
    lac_ank_add2 double precision,
    lac_ank_add3 double precision,
    lac_aerk_vo2mlkgmin double precision,
    lac_aerk_vo2mlmin double precision,
    lac_aerk_vo2 double precision,
    lac_aerk_vo2pr double precision,
    lac_aerk_hr double precision,
    lac_aerk_speed double precision,
    lac_aerk_pace double precision,
    lac_aerk_p double precision,
    lac_aerk_pkg double precision,
    lac_aerk_angle double precision,
    lac_aerk_lac double precision,
    lac_aerk_add1 double precision,
    lac_aerk_add2 double precision,
    lac_aerk_add3 double precision,
    vent_ank_vo2mlkgmin double precision,
    vent_ank_vo2mlmin double precision,
    vent_ank_vo2 double precision,
    vent_ank_vo2pr double precision,
    vent_ank_hr double precision,
    vent_ank_speed double precision,
    vent_ank_pace double precision,
    vent_ank_p double precision,
    vent_ank_pkg double precision,
    vent_ank_angle double precision,
    vent_ank_lac double precision,
    vent_ank_add1 double precision,
    vent_ank_add2 double precision,
    vent_ank_add3 double precision,
    vent_aerk_vo2mlkgmin double precision,
    vent_aerk_vo2mlmin double precision,
    vent_aerk_vo2 double precision,
    vent_aerk_vo2pr double precision,
    vent_aerk_hr double precision,
    vent_aerk_speed double precision,
    vent_aerk_pace double precision,
    vent_aerk_p double precision,
    vent_aerk_pkg double precision,
    vent_aerk_angle double precision,
    vent_aerk_lac double precision,
    vent_aerk_add1 double precision,
    vent_aerk_add2 double precision,
    vent_aerk_add3 double precision,
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint
);


ALTER TABLE public.dirresults OWNER TO klabadmin;

--
-- Name: dirresults_iddirresults_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.dirresults_iddirresults_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dirresults_iddirresults_seq OWNER TO klabadmin;

--
-- Name: dirresults_iddirresults_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.dirresults_iddirresults_seq OWNED BY public.dirresults.iddirresults;


--
-- Name: dirtest; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.dirtest (
    iddirtest integer NOT NULL,
    idmeasurement integer NOT NULL,
    meascols text,
    weightkg double precision,
    heightcm double precision,
    bmi double precision,
    fat_pr double precision,
    fat_p1 double precision,
    fat_p2 double precision,
    fat_p3 double precision,
    fat_p4 double precision,
    fat_style integer,
    fat_equip character varying(45),
    fvc double precision,
    fev1 double precision,
    air_press double precision,
    air_temp double precision,
    air_humid double precision,
    testprotocol character varying(45),
    air_press_unit integer,
    settingslist text,
    lt1_x double precision,
    lt1_y double precision,
    lt2_x double precision,
    lt2_y double precision,
    vt1_x double precision,
    vt2_x double precision,
    vt1_y double precision,
    vt2_y double precision,
    lt1_calc_x double precision,
    lt1_calc_y double precision,
    lt2_calc_x double precision,
    lt2_calc_y double precision,
    protocolmodel smallint,
    testtype smallint,
    protocolxval smallint,
    steptime integer,
    w_rest smallint,
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint,
    norawdata smallint
);


ALTER TABLE public.dirtest OWNER TO klabadmin;

--
-- Name: dirtest_iddirtest_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.dirtest_iddirtest_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dirtest_iddirtest_seq OWNER TO klabadmin;

--
-- Name: dirtest_iddirtest_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.dirtest_iddirtest_seq OWNED BY public.dirtest.iddirtest;


--
-- Name: dirteststeps; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.dirteststeps (
    iddirteststeps integer NOT NULL,
    idmeasurement integer NOT NULL,
    stepno integer,
    ana_time integer,
    timestop double precision,
    speed double precision,
    pace double precision,
    angle double precision,
    elev double precision,
    vo2calc double precision,
    t_tot double precision,
    t_ex double precision,
    fico2 double precision,
    fio2 double precision,
    feco2 double precision,
    feo2 double precision,
    vde double precision,
    vco2 double precision,
    vo2 double precision,
    bf double precision,
    ve double precision,
    petco2 double precision,
    peto2 double precision,
    vo2kg double precision,
    re double precision,
    hr double precision,
    la double precision,
    rer double precision,
    ve_stpd double precision,
    veo2 double precision,
    veco2 double precision,
    tv double precision,
    ee_ae double precision,
    la_vo2 double precision,
    o2pulse double precision,
    vde_tv double precision,
    va double precision,
    o2sa double precision,
    rpe double precision,
    bp_sys double precision,
    bp_dia double precision,
    own1 double precision,
    own2 double precision,
    own3 double precision,
    own4 double precision,
    own5 double precision,
    step_is_rest integer,
    step_is_30max integer,
    step_is_60max integer,
    step_is_rec integer,
    calc_start integer,
    calc_end integer,
    comments character varying(100),
    timestart double precision,
    duration double precision,
    eco double precision,
    p double precision,
    wkg double precision,
    vo2_30s double precision,
    vo2_pr double precision,
    step_is_last integer,
    deleted smallint,
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    created_date timestamp without time zone,
    modded smallint,
    own6 double precision,
    own7 double precision,
    own8 double precision,
    own9 double precision,
    own10 double precision,
    to2 double precision,
    tco2 double precision
);


ALTER TABLE public.dirteststeps OWNER TO klabadmin;

--
-- Name: dirteststeps_iddirteststeps_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.dirteststeps_iddirteststeps_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dirteststeps_iddirteststeps_seq OWNER TO klabadmin;

--
-- Name: dirteststeps_iddirteststeps_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.dirteststeps_iddirteststeps_seq OWNED BY public.dirteststeps.iddirteststeps;


--
-- Name: measurement_list; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.measurement_list (
    idmeasurement integer NOT NULL,
    measname character varying(45),
    idcustomer integer NOT NULL,
    tablename text,
    idpatterndef text,
    do_year smallint,
    do_month smallint,
    do_day smallint,
    do_hour smallint,
    do_min smallint,
    sessionno integer,
    info text,
    measurements text,
    groupnotes text,
    cbcharts text,
    cbcomments text,
    created_by bigint,
    mod_by bigint,
    mod_date timestamp without time zone,
    deleted smallint,
    created_date timestamp without time zone,
    modded smallint,
    test_location character varying(65),
    keywords text,
    tester_name character varying(65),
    modder_name character varying(65),
    meastype integer,
    sent_to_sprintai timestamp without time zone
);


ALTER TABLE public.measurement_list OWNER TO klabadmin;

--
-- Name: measurement_list_idmeasurement_seq; Type: SEQUENCE; Schema: public; Owner: klabadmin
--

CREATE SEQUENCE public.measurement_list_idmeasurement_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.measurement_list_idmeasurement_seq OWNER TO klabadmin;

--
-- Name: measurement_list_idmeasurement_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: klabadmin
--

ALTER SEQUENCE public.measurement_list_idmeasurement_seq OWNED BY public.measurement_list.idmeasurement;


--
-- Name: sportti_id_list; Type: TABLE; Schema: public; Owner: klabadmin
--

CREATE TABLE public.sportti_id_list (
    sportti_id character varying(64) NOT NULL
);


ALTER TABLE public.sportti_id_list OWNER TO klabadmin;

--
-- Name: customer idcustomer; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.customer ALTER COLUMN idcustomer SET DEFAULT nextval('public.customer_idcustomer_seq'::regclass);


--
-- Name: dirrawdata iddirrawdata; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirrawdata ALTER COLUMN iddirrawdata SET DEFAULT nextval('public.dirrawdata_iddirrawdata_seq'::regclass);


--
-- Name: dirreport iddirreport; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirreport ALTER COLUMN iddirreport SET DEFAULT nextval('public.dirreport_iddirreport_seq'::regclass);


--
-- Name: dirresults iddirresults; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirresults ALTER COLUMN iddirresults SET DEFAULT nextval('public.dirresults_iddirresults_seq'::regclass);


--
-- Name: dirtest iddirtest; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirtest ALTER COLUMN iddirtest SET DEFAULT nextval('public.dirtest_iddirtest_seq'::regclass);


--
-- Name: dirteststeps iddirteststeps; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirteststeps ALTER COLUMN iddirteststeps SET DEFAULT nextval('public.dirteststeps_iddirteststeps_seq'::regclass);


--
-- Name: measurement_list idmeasurement; Type: DEFAULT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.measurement_list ALTER COLUMN idmeasurement SET DEFAULT nextval('public.measurement_list_idmeasurement_seq'::regclass);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (idcustomer);


--
-- Name: customer customer_sportti_id_key; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_sportti_id_key UNIQUE (sportti_id);


--
-- Name: dirrawdata dirrawdata_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirrawdata
    ADD CONSTRAINT dirrawdata_pkey PRIMARY KEY (iddirrawdata);


--
-- Name: dirreport dirreport_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirreport
    ADD CONSTRAINT dirreport_pkey PRIMARY KEY (iddirreport);


--
-- Name: dirresults dirresults_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirresults
    ADD CONSTRAINT dirresults_pkey PRIMARY KEY (iddirresults);


--
-- Name: dirtest dirtest_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirtest
    ADD CONSTRAINT dirtest_pkey PRIMARY KEY (iddirtest);


--
-- Name: dirteststeps dirteststeps_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirteststeps
    ADD CONSTRAINT dirteststeps_pkey PRIMARY KEY (iddirteststeps);


--
-- Name: measurement_list measurement_list_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.measurement_list
    ADD CONSTRAINT measurement_list_pkey PRIMARY KEY (idmeasurement);


--
-- Name: sportti_id_list sportti_id_list_pkey; Type: CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.sportti_id_list
    ADD CONSTRAINT sportti_id_list_pkey PRIMARY KEY (sportti_id);


--
-- Name: dirrawdata dirrawdata_idmeasurement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirrawdata
    ADD CONSTRAINT dirrawdata_idmeasurement_fkey FOREIGN KEY (idmeasurement) REFERENCES public.measurement_list(idmeasurement) ON DELETE CASCADE;


--
-- Name: dirreport dirreport_idmeasurement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirreport
    ADD CONSTRAINT dirreport_idmeasurement_fkey FOREIGN KEY (idmeasurement) REFERENCES public.measurement_list(idmeasurement) ON DELETE CASCADE;


--
-- Name: dirresults dirresults_idmeasurement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirresults
    ADD CONSTRAINT dirresults_idmeasurement_fkey FOREIGN KEY (idmeasurement) REFERENCES public.measurement_list(idmeasurement) ON DELETE CASCADE;


--
-- Name: dirtest dirtest_idmeasurement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirtest
    ADD CONSTRAINT dirtest_idmeasurement_fkey FOREIGN KEY (idmeasurement) REFERENCES public.measurement_list(idmeasurement) ON DELETE CASCADE;


--
-- Name: dirteststeps dirteststeps_idmeasurement_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.dirteststeps
    ADD CONSTRAINT dirteststeps_idmeasurement_fkey FOREIGN KEY (idmeasurement) REFERENCES public.measurement_list(idmeasurement) ON DELETE CASCADE;


--
-- Name: measurement_list measurement_list_idcustomer_fkey; Type: FK CONSTRAINT; Schema: public; Owner: klabadmin
--

ALTER TABLE ONLY public.measurement_list
    ADD CONSTRAINT measurement_list_idcustomer_fkey FOREIGN KEY (idcustomer) REFERENCES public.customer(idcustomer) ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO klabadmin;


--
-- PostgreSQL database dump complete
--

