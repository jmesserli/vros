# vros - Verlesen Replacing Overengineered Solution

## Redis Data Structures

### Cards

A card is a hash. The key is its serial number, prefixed with "card:". Fields contained:

- name: The full name of the owner
- email: The email of the owner
- register_code: The temporary register code. Is nonexistent if the user is registered

### Stamp

The stamps of one day are aggregated in a hash. The key follows the schema "stamps:yyyy-mm-dd".

One stamp is represented by a key / value pair, where the key is the cards serial number and the value is a comma-separated string of times
