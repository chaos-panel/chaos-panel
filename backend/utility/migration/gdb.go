package migration

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// get db group and nodes from config
func GetDatabaseGroups() map[string]*gdb.ConfigNode {
	var (
		ctx           = context.Background()
		configMap     map[string]interface{}
		configNodeKey = "database"
	)

	if v, _ := gcfg.Instance().Get(ctx, configNodeKey); !v.IsEmpty() {
		configMap = v.Map()
	}

	nodes := make(map[string]*gdb.ConfigNode)

	if len(configMap) == 0 {
		return nodes
	}

	for key, val := range configMap {
		var node *gdb.ConfigNode
		switch value := val.(type) {
		case []interface{}:
			for _, v := range value {
				if node = parseDBConfigNode(v); node != nil {
					break
				}
			}
		case map[string]interface{}:
			node = parseDBConfigNode(value)
		}
		parsedLinkNode, err := parseConfigNodeLink(node)
		if err == nil {
			nodes[key] = parsedLinkNode
		}
	}

	return nodes
}

func parseDBConfigNode(value interface{}) *gdb.ConfigNode {
	nodeMap, ok := value.(map[string]interface{})
	if !ok {
		panic(errors.New("config invalid"))
	}
	var (
		node = &gdb.ConfigNode{}
		err  = gconv.Struct(nodeMap, node)
	)
	if err != nil {
		panic(err)
	}
	// Find possible `Link` configuration content.
	if _, v := gutil.MapPossibleItemByKey(nodeMap, "Link"); v != nil {
		node.Link = gconv.String(v)
	}
	return node
}

func parseConfigNodeLink(node *gdb.ConfigNode) (*gdb.ConfigNode, error) {
	var (
		defaultProtocol        = `tcp`
		linkPattern            = `^(\w+):(.*?):(.*?)@(\w+?)\((.+?)\)/{0,1}([^\?]*)\?{0,1}(.*?)$`
		linkPatternDescription = `type:username:password@protocol(host:port)/dbname?param1=value1&...&paramN=valueN`
		defaultCharset         = `utf8`
		link                   = node.Link
		match                  []string
	)
	if link != "" {
		// To be compatible with old configuration,
		// it checks and converts the link to new configuration.
		if node.Type != "" && !gstr.HasPrefix(link, node.Type+":") {
			link = fmt.Sprintf("%s:%s", node.Type, link)
		}
		match, _ = gregex.MatchString(linkPattern, link)
		if len(match) <= 5 {
			return nil, gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid link configuration: %s, shuold be pattern like: %s`,
				link, linkPatternDescription,
			)
		}
		node.Type = match[1]
		node.User = match[2]
		node.Pass = match[3]
		node.Protocol = match[4]
		array := gstr.Split(match[5], ":")
		if node.Protocol == "file" {
			node.Name = match[5]
		} else {
			if len(array) == 2 {
				// link with port.
				node.Host = array[0]
				node.Port = array[1]
			} else {
				// link without port.
				node.Host = array[0]
			}
			node.Name = match[6]
		}
		if len(match) > 6 && match[7] != "" {
			node.Extra = match[7]
		}
	}
	if node.Extra != "" {
		if m, _ := gstr.Parse(node.Extra); len(m) > 0 {
			_ = gconv.Struct(m, &node)
		}
	}
	// Default value checks.
	if node.Charset == "" {
		node.Charset = defaultCharset
	}
	if node.Protocol == "" {
		node.Protocol = defaultProtocol
	}
	return node, nil
}
