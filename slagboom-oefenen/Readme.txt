Beschrijving van de code:

Dit Go-programma begroet de gebruiker op basis van het tijdstip van de dag en biedt vervolgens verschillende opties:
Kenteken Registreren: De gebruiker kan een nieuw kenteken registreren door de gebruikersnaam en het kenteken in te voeren. Deze informatie wordt opgeslagen in een JSON-bestand genaamd "bookings.json".
Toegang Park Controleren: De gebruiker kan controleren of zijn/haar kenteken toegang geeft tot het park door het kenteken in te voeren. Het systeem controleert vervolgens of het kenteken geldig is aan de hand van de gegevens in het "bookings.json" bestand.
Gebruiker Verwijderen: De gebruiker kan een bestaande gebruiker verwijderen door het kenteken van de gebruiker in te voeren. Voordat de gebruiker wordt verwijderd, wordt een bevestigingsbericht weergegeven om te verifiëren of de gebruiker daadwerkelijk moet worden verwijderd. Als de gebruiker bevestigt, wordt deze uit het bestand verwijderd.
Gebruiker Status Wijzigen: De gebruiker kan de status van een bestaande gebruiker wijzigen door het kenteken van de gebruiker in te voeren en aan te geven of de gebruiker geactiveerd moet worden of niet.
Exit: De gebruiker kan het programma verlaten door deze optie te kiezen.

Exitcodes en logging:

Als het programma het JSON-bestand niet kan openen, decoderen of schrijven, wordt een foutlogboek gegenereerd en het programma stopt met exitcode 1.
Als het logbestand niet kan worden geopend, wordt een foutbericht afgedrukt op de standaarduitvoer en het programma stopt met exitcode 1.
Voor de verwijderingsfunctie wordt een bevestigingsbericht weergegeven voordat een gebruiker wordt verwijderd. Als de gebruiker "ja" antwoordt, wordt de gebruiker verwijderd en wordt een bevestigingsbericht weergegeven. Als de gebruiker "nee" antwoordt, wordt de verwijderingsactie geannuleerd.
Logging:

Het programma gebruikt een aangepaste logger met de naam "INFO:". Deze logger logt de datum, tijd en het bestandsnaam-regelnummer van waaruit de log wordt gemaakt. De logberichten worden geschreven naar het bestand "trace.log".