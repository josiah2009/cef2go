
release_path = ENV['RELEASE_PATH']
attempts     = (ARGV.shift || 100).to_i
timeout      = (ARGV.shift || 6).to_i
failures     = 0
successes    = 0
i            = 0
bin          = "cef2go"

def is_running?(bin, timeout)
  tries = 0
  alive = true
  while alive && tries < timeout
    sleep 1
    tries += 1
    alive = `pidof #{bin}`.strip.split(' ').count > 0
  end
  alive 
end

def kill(bin)
  `killall -9 #{bin} > /dev/null 2>&1`
end

def run(release_path, bin)
  `RELEASE_PATH=#{File.expand_path(release_path)} ./#{bin} >> test.log 2>&1 &`
end

puts "Running #{bin} #{attempts} times"

while i < attempts
  run(release_path, bin)
  if is_running?(bin, timeout)
    successes += 1 
    print '.'
  else
    failures += 1
    print 'F'
  end
  kill(bin)
  i+=1
end

puts

puts "Successes: #{successes}. Failures: #{failures}. Fail Rate: #{(failures.to_f / attempts.to_f) * 100}"
