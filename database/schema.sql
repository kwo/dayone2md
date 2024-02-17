CREATE TABLE ZATTACHMENT ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZFAVORITE INTEGER, ZFILESIZE INTEGER, ZHASDATA INTEGER, ZHEIGHT INTEGER, ZORDERINENTRY INTEGER, ZWIDTH INTEGER, ZENTRY INTEGER, ZLOCATION INTEGER, ZTHUMBNAIL INTEGER, ZISRECORDING INTEGER, ZISSKETCH INTEGER, ZISO INTEGER, ZBOOKBACKCOVER INTEGER, ZBOOKFRONTCOVER INTEGER, Z_FOK_BOOKFRONTCOVER INTEGER, Z_FOK_BOOKBACKCOVER INTEGER, ZDATE TIMESTAMP, ZDURATION FLOAT, ZLASTAPPLEIDENTIFIERUPDATEERRORDATE TIMESTAMP, ZLASTMEDIAMATCHATTEMPTDATE TIMESTAMP, ZUSERMODIFIEDDATE TIMESTAMP, ZEXPOSUREBIASVALUE FLOAT, ZAPPLECLOUDIDENTIFIER VARCHAR, ZAPPLELOCALIDENTIFIER VARCHAR, ZCREATIONDEVICE VARCHAR, ZCREATIONDEVICEIDENTIFIER VARCHAR, ZDO_ENTITYNAME VARCHAR, ZIDENTIFIER VARCHAR, ZMD5 VARCHAR, ZRECORDINGDEVICE VARCHAR, ZTYPE VARCHAR, ZAUDIOCHANNELS VARCHAR, ZFORMAT VARCHAR, ZSAMPLERATE VARCHAR, ZTIMEZONENAME VARCHAR, ZTITLE VARCHAR, ZTRANSCRIPTION VARCHAR, ZPDFNAME VARCHAR, ZCAMERAMAKE VARCHAR, ZCAMERAMODEL VARCHAR, ZCAPTION VARCHAR, ZFILENAME VARCHAR, ZFNUMBER VARCHAR, ZFOCALLENGTH VARCHAR, ZLENSMAKE VARCHAR, ZLENSMODEL VARCHAR, ZSOURCEIDENTIFIER VARCHAR , ZEMBED INTEGER);
CREATE TABLE Z_14TAGS ( Z_14ENTRIES INTEGER, Z_55TAGS1 INTEGER, PRIMARY KEY (Z_14ENTRIES, Z_55TAGS1) );
CREATE TABLE ZJOURNAL ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZADDLOCATIONTONEWENTRIES INTEGER, ZCOLORHEX INTEGER, ZCOMMENTSENABLED INTEGER, ZCONCEAL INTEGER, ZHASCHECKEDFORREMOTEJOURNAL INTEGER, ZHIDDEN INTEGER, ZIMPORTING INTEGER, ZISTRASHJOURNAL INTEGER, ZLOCALONLY INTEGER, ZPLACEHOLDERFORENCRYPTEDJOURNAL INTEGER, ZREADONLY INTEGER, ZSHOULDBEINCLUDEDINONTHISDAY INTEGER, ZSHOULDBEINCLUDEDINSTREAKS INTEGER, ZSHOULDBEINCLUDEDINTODAYVIEW INTEGER, ZSHOULDROTATEKEYS INTEGER, ZSHOULDSUPPRESSUSERPUSHNOTIFICATIONS INTEGER, ZSORTORDER INTEGER, ZTYPE INTEGER, ZWANTSENCRYPTION INTEGER, ZSHAREDJOURNALINFO INTEGER, ZDIRTYMODIFICATIONDATE TIMESTAMP, ZLASTTRASHEDENTRIESSYNCDATE TIMESTAMP, ZRESTRICTEDJOURNALEXPIRATIONDATE TIMESTAMP, ZACTIVEKEYFINGERPRINT VARCHAR, ZJOURNALDESCRIPTION VARCHAR, ZNAME VARCHAR, ZOWNERID VARCHAR, ZSHAREPERMISSIONS VARCHAR, ZSORTMODE VARCHAR, ZSYNCJOURNALID VARCHAR, ZSYNCUPLOADBASEURL VARCHAR, ZTEMPLATEID VARCHAR, ZUUIDFORAUXILIARYSYNC VARCHAR, ZCONNECTEDSERVICES BLOB, ZCROPPEDCOVERIMAGEDATA BLOB, ZORIGINALCOVERIMAGEDATA BLOB, ZVAULTKEY BLOB );
CREATE TABLE ZTAG ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZNAME VARCHAR, ZNORMALIZEDNAME VARCHAR );
CREATE TABLE ZWEATHER ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZENTRY INTEGER, ZMOONPHASE FLOAT, ZPRESSUREMB FLOAT, ZRELATIVEHUMIDITY FLOAT, ZSUNRISEDATE TIMESTAMP, ZSUNSETDATE TIMESTAMP, ZTEMPERATURECELSIUS FLOAT, ZVISIBILITYKM FLOAT, ZWINDBEARING FLOAT, ZWINDCHILLCELSIUS FLOAT, ZWINDSPEEDKPH FLOAT, ZCONDITIONSDESCRIPTION VARCHAR, ZMOONPHASECODE VARCHAR, ZWEATHERCODE VARCHAR, ZWEATHERSERVICENAME VARCHAR );
CREATE TABLE ZLOCATION ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZFLOORLEVEL INTEGER, ZUSERSETMANUALLY INTEGER, ZATTACHMENT INTEGER, Z2_ATTACHMENT INTEGER, ZALTITUDE FLOAT, ZHEADING FLOAT, ZLATITUDE FLOAT, ZLONGITUDE FLOAT, ZSPEED FLOAT, ZADDRESS VARCHAR, ZADMINISTRATIVEAREA VARCHAR, ZCOUNTRY VARCHAR, ZLOCALITYNAME VARCHAR, ZPLACENAME VARCHAR, ZTIMEZONENAME VARCHAR, ZUSERLABEL VARCHAR, ZUSERTYPE VARCHAR, ZREGION BLOB , ZEMBED INTEGER);
CREATE TABLE ZENTRY ( Z_PK INTEGER PRIMARY KEY, Z_ENT INTEGER, Z_OPT INTEGER, ZCHECKLISTCOMPLETEDITEMS INTEGER, ZCHECKLISTTOTALITEMS INTEGER, ZDURATION INTEGER, ZGREGORIANDAY INTEGER, ZGREGORIANMONTH INTEGER, ZGREGORIANYEAR INTEGER, ZISALLDAY INTEGER, ZISDRAFT INTEGER, ZISPINNED INTEGER, ZSTARRED INTEGER, ZBOOK INTEGER, ZJOURNAL INTEGER, ZLOCATION INTEGER, ZMUSIC INTEGER, ZREMOTEENTRY INTEGER, ZRESTORATIONJOURNAL INTEGER, ZTEMPLATE INTEGER, ZUSERACTIVITY INTEGER, ZVISIT INTEGER, ZWEATHER INTEGER, ZCREATIONDATE TIMESTAMP, ZEDITINGTIME FLOAT, ZMODIFIEDDATE TIMESTAMP, ZNORMALIZEDGMTDATE TIMESTAMP, ZCHANGEID VARCHAR, ZCREATIONDEVICE VARCHAR, ZCREATIONDEVICEMODEL VARCHAR, ZCREATIONDEVICETYPE VARCHAR, ZCREATIONOSNAME VARCHAR, ZCREATIONOSVERSION VARCHAR, ZENTRYTYPE VARCHAR, ZFEATUREFLAGSSTRING VARCHAR, ZGREGORIANSECTIONKEY VARCHAR, ZLASTEDITINGDEVICEID VARCHAR, ZLASTEDITINGDEVICENAME VARCHAR, ZMARKDOWNTEXT VARCHAR, ZPROMPTID VARCHAR, ZRICHTEXTJSON VARCHAR, ZSOURCESTRING VARCHAR, ZUNREADMARKERID VARCHAR, ZUUID VARCHAR, ZCREATOR BLOB, ZPUBLISHURL BLOB, ZTIMEZONE BLOB , ZCOMMENTSDISABLED INTEGER, ZCOMMENTSNOTIFICATIONSDISABLED INTEGER, ZOPTOUTOFAUTOMATICSYNCUNTILDATE TIMESTAMP);
