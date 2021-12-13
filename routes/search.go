package routes

import (
	"github.com/lor00x/goldap/message"
	javaser "github.com/rakuten-tech/jndi-ldap-test-server/java/serialization"
	"github.com/rakuten-tech/jndi-ldap-test-server/util/must"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vjeantet/ldapserver"
)

// This payload is the serialized Java String "!!! VULNERABLE !!!"
// The payload was generated using the code in extra/generate-payload.kt
var vulnerableStringPayload = must.DecodeBase64("rO0ABXQAEiEhISBWVUxORVJBQkxFICEhIQ==")

func SetVulnerablePayload(payload string) {
	vulnerableStringPayload = javaser.EncodeString(payload)
}

func handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	r := m.GetSearchRequest()

	log.Info().
		Str("component", "ldap").
		Str("event", "request").
		Str("client_ip", m.Client.Addr().String()).
		Dict("request", zerolog.Dict().
			Str("type", "search").
			Str("base_dn", string(r.BaseObject())).
			Str("filter", r.FilterString()).
			Array("attributes", arrayOfLdapStrings(r.Attributes())).
			Int("time_limit", r.TimeLimit().Int()),
		).
		Msg("Incoming LDAP Search Request")

	e := ldapserver.NewSearchResultEntry("")
	e.AddAttribute("javaClassName", "foo")
	e.AddAttribute("javaSerializedData", message.AttributeValue(vulnerableStringPayload))
	w.Write(e)

	res := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess)
	w.Write(res)
}

func arrayOfLdapStrings(ldapStrings []message.LDAPString) *zerolog.Array {
	arr := zerolog.Arr()
	for _, s := range ldapStrings {
		arr.Str(string(s))
	}
	return arr
}
