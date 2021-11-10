/* RUN the sql script one by one */

/* RUN FIRST */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/* RUN SECOND */
CREATE FUNCTION trigger_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
	INSERT INTO verification
         VALUES(uuid_generate_v4(),NEW.id,(SELECT array_to_string(ARRAY(SELECT chr((97 + round(random() * 25)) :: integer) 
	FROM generate_series(1,4)), '')),(current_timestamp + interval '1 day'));
 
    RETURN NEW;
END;
$$

/* RUN THIRD */
CREATE TRIGGER trigger_verification
  AFTER INSERT
  ON credential
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_function();