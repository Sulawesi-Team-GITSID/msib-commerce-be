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

/* RUN FOURTH */
CREATE FUNCTION seller_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
	update credential set seller = true where id = new.credential_id;
    RETURN NEW;
END;
$$

/* RUN Fifth */
CREATE TRIGGER trigger_seller
  AFTER INSERT
  ON shop
  FOR EACH ROW
  EXECUTE PROCEDURE seller_function();

/* for sixth */
CREATE OR REPLACE FUNCTION public.file_function()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
        IF NEW.path = 'profile' THEN
            update public.profile set image_url = new.image_url where id = new.entity_id;
        END IF;
        IF NEW.path = 'game' THEN
            update public.game set game.image_url = new.image_url where profile.id = new.entity_id;
        END IF;
    RETURN NEW;
END;
$function$
;



CREATE TRIGGER file_activate
  AFTER INSERT OR UPDATE OR DELETE
  ON file
  FOR EACH ROW
  EXECUTE PROCEDURE file_function();
