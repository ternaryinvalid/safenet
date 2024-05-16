CREATE SCHEMA messages;

ALTER SCHEMA messages OWNER TO user_admin;

CREATE TABLE IF NOT EXISTS messages.messages (
    message_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
    public_key_from text NOT NULL,
    public_key_to text NOT NULL,
    message_data text NOT NULL,
    dt timestamp WITHOUT TIME ZONE DEFAULT NOW() NOT NULL
);

ALTER TABLE messages.messages OWNER TO user_admin;

--------------------------------------------------------------------------------

ALTER TABLE messages.messages
    ADD CONSTRAINT pk_messages_messages PRIMARY KEY (message_id);

--------------------------------------------------------------------------------
DROP FUNCTION IF EXISTS messages.messages_create(_public_key_from text, _public_key_to text, _message text);
CREATE FUNCTION messages.message_create(_public_key_from text, _public_key_to text, _message text) RETURNS TABLE(message_id bigint)
    LANGUAGE plpgsql SECURITY DEFINER
AS $$
BEGIN
    RETURN QUERY (
    WITH insert_messages_cte AS (
        INSERT INTO messages.messages AS m
            (public_key_from, public_key_to, message_data)
        VALUES (_public_key_from, _public_key_to, _message)
        RETURNING m.message_id)
            SELECT m.message_id
            FROM insert_messages_cte m
    );
END;
$$;

ALTER FUNCTION messages.message_create(text, text, text) OWNER TO user_admin;

--------------------------------------------------------------------------------

DROP FUNCTION IF EXISTS messages.messages_get(_public_key_to text);
CREATE FUNCTION messages.messages_get(_public_key_to text)
    RETURNS TABLE (public_key_from text, message_data text, dt timestamp)
    LANGUAGE plpgsql SECURITY DEFINER
AS $$
BEGIN
    RETURN QUERY
        SELECT
            m.public_key_from,
            m.message_data,
            m.dt
        FROM messages.messages m
        WHERE m.public_key_to = _public_key_to;
END;
$$;

ALTER FUNCTION messages.messages_get(text) OWNER TO user_admin;

--------------------------------------------------------------------------------
