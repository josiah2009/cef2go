release_path = ENV['RELEASE_PATH']
attempts     = (ARGV.shift || 100).to_i
timeout      = (ARGV.shift || 6).to_i
failures     = 0
successes    = 0
i            = 0
bin          = "cef2go"
args         = "--disable-gpu -v DEBUG --single-process --disable-gpu --disable-gpu-compositing"

def kill(bin)
  `killall -9 #{bin} > /dev/null 2>&1`
end

def run(release_path, bin, args)
  `echo "############" >> test.log`
  `RELEASE_PATH=#{File.expand_path(release_path)} ./#{bin} #{args} >> test.log 2>&1`
  $?.to_i == 0
end

puts "Running #{bin} #{attempts} times"

while i < attempts
  if run(release_path, bin, args)
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
