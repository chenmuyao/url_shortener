-- Random URL Generator Function
local function random_string(length)
	local chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	local result = {}
	for _ = 1, length do
		local index = math.random(1, #chars)
		table.insert(result, chars:sub(index, index))
	end
	return table.concat(result)
end

local function random_domain()
	local tlds = { "com", "org", "net", "io", "xyz", "info", "biz" }
	local subdomain_chance = math.random()
	local subdomain = (subdomain_chance > 0.7) and (random_string(math.random(3, 8)) .. ".") or ""
	local domain_name = random_string(math.random(5, 10))
	local tld = tlds[math.random(1, #tlds)]
	return subdomain .. domain_name .. "." .. tld
end

local function random_path()
	local path_segments = math.random(1, 5)
	local path = {}
	for _ = 1, path_segments do
		table.insert(path, random_string(math.random(3, 10)))
	end
	return "/" .. table.concat(path, "/")
end

local function random_query_params()
	local param_count = math.random(1, 4)
	local params = {}
	for _ = 1, param_count do
		local key = random_string(math.random(3, 8))
		local value = random_string(math.random(3, 8))
		table.insert(params, key .. "=" .. value)
	end
	return "?" .. table.concat(params, "&")
end

local function generate_url()
	local schemes = { "http", "https" }
	local scheme = schemes[math.random(1, #schemes)]
	local domain = random_domain()
	local path = (math.random() > 0.5) and random_path() or ""
	local query = (math.random() > 0.7) and random_query_params() or ""
	return scheme .. "://" .. domain .. path .. query
end

-- Setup function for wrk
request = function()
	-- Generate a random URL
	local url = generate_url()

	-- Prepare the request
	wrk.method = "POST"
	wrk.body = '{"url": "' .. url .. '"}'
	wrk.headers["Content-Type"] = "application/json"

	return wrk.format(nil, wrk.path)
end
