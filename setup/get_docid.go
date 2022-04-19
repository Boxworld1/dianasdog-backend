package setup

import (
	"github.com/beevik/etree"
)

func GetDocid(item *etree.Element, targetResource string) string {
	key := item.SelectElement("key")
	docid := targetResource + "@" + key.Text()
	return docid
}
