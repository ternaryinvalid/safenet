CREATE SCHEMA messages;

ALTER SCHEMA messages OWNER TO user_admin;

CREATE TABLE IF NOT EXISTS messages.messages (
    message_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
    public_key_from varchar NOT NULL,
    public_key_to varchar NOT NULL,
    message_data varchar NOT NULL,
    dt timestamp WITHOUT TIME ZONE DEFAULT NOW() NOT NULL
);

ALTER TABLE messages.messages OWNER TO user_admin;

--------------------------------------------------------------------------------

ALTER TABLE messages.messages
    ADD CONSTRAINT pk_messages_messages PRIMARY KEY (message_id);

--------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION messages.message_create(_public_key_from varchar, _public_key_to varchar, _message varchar) RETURNS TABLE(message_id bigint)
    LANGUAGE plpgsql SECURITY DEFINER
AS $$
BEGIN
    WITH insert_messages_cte AS (
        INSERT INTO messages.messages AS m
            (public_key_from, public_key_to, message_data)
            VALUES (_public_key_from, _public_key_to, _message)
            RETURNING m.message_id
    )
    SELECT m.message_id
    FROM insert_messages_cte m;
END;
$$;

ALTER FUNCTION messages.message_create(_public_key_from varchar, _public_key_to varchar, _message varchar) OWNER TO user_admin;

--------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION messages.messages_get(_public_key_to varchar, _limit integer DEFAULT 50)
    RETURNS TABLE (public_key_from varchar, message_data varchar)
    LANGUAGE plpgsql SECURITY DEFINER
AS $$
BEGIN
    RETURN QUERY
        SELECT
            m.public_key_from,
            m.message_data
        FROM messages.messages m
        WHERE m.public_key_to = _public_key_to
        LIMIT _limit;
END;
$$;

ALTER FUNCTION messages.messages_get(_public_key_to varchar, _limit integer) OWNER TO user_admin;
