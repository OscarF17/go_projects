--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6
-- Dumped by pg_dump version 15.6

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
-- Name: create_producto(text, double precision); Type: PROCEDURE; Schema: public; Owner: oscar
--

CREATE PROCEDURE public.create_producto(IN inombre text, IN iprecio double precision)
    LANGUAGE plpgsql
    AS $$
begin
insert into productos (nombre, precio) values (inombre, iprecio);
end;
$$;


ALTER PROCEDURE public.create_producto(IN inombre text, IN iprecio double precision) OWNER TO oscar;

--
-- Name: delete_producto(integer); Type: PROCEDURE; Schema: public; Owner: oscar
--

CREATE PROCEDURE public.delete_producto(IN iid integer)
    LANGUAGE plpgsql
    AS $$
begin
delete from productos where id = iid;
end;
$$;


ALTER PROCEDURE public.delete_producto(IN iid integer) OWNER TO oscar;

--
-- Name: get_producto(integer); Type: FUNCTION; Schema: public; Owner: oscar
--

CREATE FUNCTION public.get_producto(id_prod integer) RETURNS TABLE(id integer, nombre text, precio double precision)
    LANGUAGE plpgsql
    AS $$
begin
return query select * from productos p where p.id = id_prod;
end;
$$;


ALTER FUNCTION public.get_producto(id_prod integer) OWNER TO oscar;

--
-- Name: update_producto(integer, text, double precision); Type: PROCEDURE; Schema: public; Owner: oscar
--

CREATE PROCEDURE public.update_producto(IN iid integer, IN inombre text, IN iprecio double precision)
    LANGUAGE plpgsql
    AS $$
begin
update productos set nombre = inombre, precio = iprecio where id = iid;
end;
$$;


ALTER PROCEDURE public.update_producto(IN iid integer, IN inombre text, IN iprecio double precision) OWNER TO oscar;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: clientes; Type: TABLE; Schema: public; Owner: oscar
--

CREATE TABLE public.clientes (
    id integer NOT NULL
);


ALTER TABLE public.clientes OWNER TO oscar;

--
-- Name: clientes_id_seq; Type: SEQUENCE; Schema: public; Owner: oscar
--

CREATE SEQUENCE public.clientes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clientes_id_seq OWNER TO oscar;

--
-- Name: clientes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: oscar
--

ALTER SEQUENCE public.clientes_id_seq OWNED BY public.clientes.id;


--
-- Name: productos; Type: TABLE; Schema: public; Owner: oscar
--

CREATE TABLE public.productos (
    id integer NOT NULL,
    nombre text,
    precio double precision
);


ALTER TABLE public.productos OWNER TO oscar;

--
-- Name: productos_id_seq; Type: SEQUENCE; Schema: public; Owner: oscar
--

CREATE SEQUENCE public.productos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.productos_id_seq OWNER TO oscar;

--
-- Name: productos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: oscar
--

ALTER SEQUENCE public.productos_id_seq OWNED BY public.productos.id;


--
-- Name: clientes id; Type: DEFAULT; Schema: public; Owner: oscar
--

ALTER TABLE ONLY public.clientes ALTER COLUMN id SET DEFAULT nextval('public.clientes_id_seq'::regclass);


--
-- Name: productos id; Type: DEFAULT; Schema: public; Owner: oscar
--

ALTER TABLE ONLY public.productos ALTER COLUMN id SET DEFAULT nextval('public.productos_id_seq'::regclass);


--
-- Data for Name: clientes; Type: TABLE DATA; Schema: public; Owner: oscar
--

COPY public.clientes (id) FROM stdin;
\.


--
-- Data for Name: productos; Type: TABLE DATA; Schema: public; Owner: oscar
--

COPY public.productos (id, nombre, precio) FROM stdin;
1	Lapiz	5.5
2	Borrador	2
3	Sacapuntas	3
4	Lonchera	150
\.


--
-- Name: clientes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: oscar
--

SELECT pg_catalog.setval('public.clientes_id_seq', 1, false);


--
-- Name: productos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: oscar
--

SELECT pg_catalog.setval('public.productos_id_seq', 8, true);


--
-- Name: clientes clientes_pkey; Type: CONSTRAINT; Schema: public; Owner: oscar
--

ALTER TABLE ONLY public.clientes
    ADD CONSTRAINT clientes_pkey PRIMARY KEY (id);


--
-- Name: productos productos_pkey; Type: CONSTRAINT; Schema: public; Owner: oscar
--

ALTER TABLE ONLY public.productos
    ADD CONSTRAINT productos_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

CREATE USER tienda_user WITH ENCRYPTED PASSWORD '666';
GRANT ALL ON SCHEMA public TO tienda_user;


--
-- Name: TABLE clientes; Type: ACL; Schema: public; Owner: oscar
--

GRANT ALL ON TABLE public.clientes TO tienda_user;


--
-- Name: TABLE productos; Type: ACL; Schema: public; Owner: oscar
--

GRANT ALL ON TABLE public.productos TO tienda_user;


--
-- Name: SEQUENCE productos_id_seq; Type: ACL; Schema: public; Owner: oscar
--

GRANT ALL ON SEQUENCE public.productos_id_seq TO tienda_user;


--
-- PostgreSQL database dump complete
--

