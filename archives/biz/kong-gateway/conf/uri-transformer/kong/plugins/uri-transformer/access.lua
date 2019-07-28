--[["

# 快速挂载开发机调试

mkdir -p /usr/local/share/lua/5.1/kong/plugins/uri-transformer

"]]

local _M = {}

function _M.execute(conf)

    -- 插件开关
    if conf.remotedebug_enable then
        require('mobdebug').start(conf.remotedebug_host)
    end

    local uri_captures = ngx.ctx.router_matches.uri_captures
    local upstream_uri = conf.uri_template

    if type(uri_captures) == "table" then

        -- 增加名字捕获组的引用
        for k, v in pairs(uri_captures) do
            if type(k) == "string" then
                upstream_uri = string.gsub(upstream_uri, string.format("{%s}", k), v)
                upstream_uri = string.gsub(upstream_uri, string.format("$%s", k), v)
            end
        end

        -- 增加索引的引用
        for i = 1, #(uri_captures) do
            upstream_uri = string.gsub(upstream_uri, string.format("{%s}", i), uri_captures[i])
            upstream_uri = string.gsub(upstream_uri, string.format("$%s", i), uri_captures[i])
        end

    end

    -- 每次都关闭调试开关
    require('mobdebug').done()

end

return _M
