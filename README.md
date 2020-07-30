# Termin berlinweit suchen

The program is designed to search for all available appointments in Berlin for flat registration in the nearest months 
on the website berlin.de. 
The URL of the main page "Anmeldung einer Wohnung", from which the search starts, must be specified in the environment variable 
TERMIN_SUCHEN_MAIN_IRL (you can also use .env for this). At the moment, the URL is - https://service.berlin.de/dienstleistung/120686/.
TERMIN_SUCHEN_CHECK_PERIOD is to set time in seconds for periodical checking(30 sec by default) 

