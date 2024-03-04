Beschrijving van de code:

Dit Go-programma begroet de gebruiker op basis van het tijdstip van de dag en controleert vervolgens 
of een opgegeven kenteken voorkomt in een lijst met boekingen die zijn opgeslagen in een JSON-bestand.
De gebruiker wordt gevraagd of ze een kenteken aan de lijst willen toevoegen, en als ze dat willen, 
wordt de nieuwe boeking toegevoegd aan de lijst en wordt de lijst bijgewerkt in het JSON-bestand.

Exitcodes en logging:

* Als het programma het JSON-bestand niet kan openen, wordt een foutlogboek gegenereerd en het programma stopt met exitcode 1.
* Als het decoderen van het JSON-bestand mislukt, wordt een foutlogboek gegenereerd en het programma stopt met exitcode 1.
* Als het schrijven naar het JSON-bestand mislukt bij het toevoegen van een nieuwe boeking, wordt een foutlogboek gegenereerd en het programma stopt met exitcode 1.
* Als het logbestand niet kan worden geopend, wordt een foutbericht afgedrukt op de standaarduitvoer en het programma stopt met exitcode 1.

Logging:

Het programma gebruikt een aangepaste logger met de naam "INFO:" en logt de datum, tijd en het bestandsnaam-regelnummer van waaruit de log wordt gemaakt.
De logberichten worden geschreven naar het bestand "testlogfile.log".