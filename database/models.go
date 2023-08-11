// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
)

type Z12TAG struct {
	Z12entries int64
	Z55tags1   int64
}

type ZATTACHMENT struct {
	ZPk                                 int64
	ZEnt                                sql.NullInt64
	ZOpt                                sql.NullInt64
	Zfavorite                           sql.NullInt64
	Zfilesize                           sql.NullInt64
	Zhasdata                            sql.NullInt64
	Zheight                             sql.NullInt64
	Zorderinentry                       sql.NullInt64
	Zwidth                              sql.NullInt64
	Zentry                              int64
	Zlocation                           sql.NullInt64
	Zthumbnail                          sql.NullInt64
	Zisrecording                        sql.NullInt64
	Zissketch                           sql.NullInt64
	Ziso                                sql.NullInt64
	Zbookbackcover                      sql.NullInt64
	Zbookfrontcover                     sql.NullInt64
	ZFokBookfrontcover                  sql.NullInt64
	ZFokBookbackcover                   sql.NullInt64
	Zdate                               sql.NullTime
	Zduration                           sql.NullFloat64
	Zlastappleidentifierupdateerrordate sql.NullTime
	Zlastmediamatchattemptdate          sql.NullTime
	Zusermodifieddate                   sql.NullTime
	Zexposurebiasvalue                  sql.NullFloat64
	Zapplecloudidentifier               sql.NullString
	Zapplelocalidentifier               sql.NullString
	Zcreationdevice                     sql.NullString
	Zcreationdeviceidentifier           sql.NullString
	ZdoEntityname                       sql.NullString
	Zidentifier                         string
	Zmd5                                string
	Zrecordingdevice                    sql.NullString
	Ztype                               string
	Zaudiochannels                      sql.NullString
	Zformat                             sql.NullString
	Zsamplerate                         sql.NullString
	Ztimezonename                       sql.NullString
	Ztitle                              sql.NullString
	Ztranscription                      sql.NullString
	Zpdfname                            sql.NullString
	Zcameramake                         sql.NullString
	Zcameramodel                        sql.NullString
	Zcaption                            sql.NullString
	Zfilename                           sql.NullString
	Zfnumber                            sql.NullString
	Zfocallength                        sql.NullString
	Zlensmake                           sql.NullString
	Zlensmodel                          sql.NullString
	Zsourceidentifier                   sql.NullString
}

type ZENTRY struct {
	ZPk                            int64
	ZEnt                           sql.NullInt64
	ZOpt                           sql.NullInt64
	Zchecklistcompleteditems       sql.NullInt64
	Zchecklisttotalitems           sql.NullInt64
	Zduration                      sql.NullInt64
	Zgregorianday                  sql.NullInt64
	Zgregorianmonth                sql.NullInt64
	Zgregorianyear                 sql.NullInt64
	Zisallday                      sql.NullInt64
	Zisdraft                       sql.NullInt64
	Zispinned                      sql.NullInt64
	Zstarred                       sql.NullInt64
	Zbook                          sql.NullInt64
	Zjournal                       int64
	Zlocation                      sql.NullInt64
	Zmusic                         sql.NullInt64
	Zremoteentry                   sql.NullInt64
	Zrestorationjournal            sql.NullInt64
	Ztemplate                      sql.NullInt64
	Zuseractivity                  sql.NullInt64
	Zvisit                         sql.NullInt64
	Zweather                       sql.NullInt64
	Zcreationdate                  string
	Zeditingtime                   sql.NullFloat64
	Zmodifieddate                  string
	Znormalizedgmtdate             sql.NullTime
	Zchangeid                      sql.NullString
	Zcreationdevice                sql.NullString
	Zcreationdevicemodel           sql.NullString
	Zcreationdevicetype            sql.NullString
	Zcreationosname                sql.NullString
	Zcreationosversion             sql.NullString
	Zentrytype                     sql.NullString
	Zfeatureflagsstring            sql.NullString
	Zgregoriansectionkey           sql.NullString
	Zlasteditingdeviceid           sql.NullString
	Zlasteditingdevicename         sql.NullString
	Zmarkdowntext                  sql.NullString
	Zpromptid                      sql.NullString
	Zrichtextjson                  sql.NullString
	Zsourcestring                  sql.NullString
	Zunreadmarkerid                sql.NullString
	Zuuid                          string
	Zcreator                       []byte
	Zpublishurl                    []byte
	Ztimezone                      []byte
	Zcommentsdisabled              sql.NullInt64
	Zcommentsnotificationsdisabled sql.NullInt64
}

type ZJOURNAL struct {
	ZPk                                  int64
	ZEnt                                 sql.NullInt64
	ZOpt                                 sql.NullInt64
	Zcolorhex                            sql.NullInt64
	Zconceal                             sql.NullInt64
	Zhascheckedforremotejournal          sql.NullInt64
	Zhidden                              sql.NullInt64
	Zimporting                           sql.NullInt64
	Zistrashjournal                      sql.NullInt64
	Zlocalonly                           sql.NullInt64
	Zplaceholderforencryptedjournal      sql.NullInt64
	Zshouldbeincludedinonthisday         sql.NullInt64
	Zshouldbeincludedinstreaks           sql.NullInt64
	Zshouldbeincludedintodayview         sql.NullInt64
	Zsortorder                           sql.NullInt64
	Zwantsencryption                     sql.NullInt64
	Zdirtymodificationdate               sql.NullTime
	Zlasttrashedentriessyncdate          sql.NullTime
	Zrestrictedjournalexpirationdate     sql.NullTime
	Zactivekeyfingerprint                sql.NullString
	Zjournaldescription                  sql.NullString
	Zname                                string
	Zownerid                             sql.NullString
	Zsharepermissions                    sql.NullString
	Zsortmode                            sql.NullString
	Zsyncjournalid                       sql.NullString
	Zsyncuploadbaseurl                   sql.NullString
	Ztemplateid                          sql.NullString
	Zuuidforauxiliarysync                sql.NullString
	Zconnectedservices                   []byte
	Zvaultkey                            []byte
	Zshouldrotatekeys                    sql.NullInt64
	Ztype                                sql.NullInt64
	Zshouldsuppressuserpushnotifications sql.NullInt64
	Zreadonly                            sql.NullInt64
	Zsharedjournalinfo                   sql.NullInt64
	Zcommentsdisabled                    sql.NullInt64
}

type ZLOCATION struct {
	ZPk                 int64
	ZEnt                sql.NullInt64
	ZOpt                sql.NullInt64
	Zfloorlevel         sql.NullInt64
	Zusersetmanually    sql.NullInt64
	Zattachment         sql.NullInt64
	Z2Attachment        sql.NullInt64
	Zaltitude           sql.NullFloat64
	Zheading            sql.NullFloat64
	Zlatitude           sql.NullFloat64
	Zlongitude          sql.NullFloat64
	Zspeed              sql.NullFloat64
	Zaddress            sql.NullString
	Zadministrativearea sql.NullString
	Zcountry            sql.NullString
	Zlocalityname       sql.NullString
	Zplacename          sql.NullString
	Ztimezonename       sql.NullString
	Zuserlabel          sql.NullString
	Zusertype           sql.NullString
	Zregion             []byte
}

type ZTAG struct {
	ZPk             int64
	ZEnt            sql.NullInt64
	ZOpt            sql.NullInt64
	Zname           string
	Znormalizedname string
}

type ZWEATHER struct {
	ZPk                    int64
	ZEnt                   sql.NullInt64
	ZOpt                   sql.NullInt64
	Zentry                 sql.NullInt64
	Zmoonphase             sql.NullFloat64
	Zpressuremb            sql.NullFloat64
	Zrelativehumidity      sql.NullFloat64
	Zsunrisedate           sql.NullString
	Zsunsetdate            sql.NullString
	Ztemperaturecelsius    sql.NullFloat64
	Zvisibilitykm          sql.NullFloat64
	Zwindbearing           sql.NullFloat64
	Zwindchillcelsius      sql.NullFloat64
	Zwindspeedkph          sql.NullFloat64
	Zconditionsdescription sql.NullString
	Zmoonphasecode         sql.NullString
	Zweathercode           sql.NullString
	Zweatherservicename    sql.NullString
}
