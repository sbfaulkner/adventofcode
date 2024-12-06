require "rake/testtask"
require "erb"
require "ostruct"

task default: "test"

Rake::TestTask.new do |t|
  t.libs << "pkg"
  t.test_files = FileList["pkg/*/day*_test.rb"]
end

desc "Generate files for a new day's puzzle (Usage: rake generate[day])"
task :generate, [:day] do |t, args|
  day = args[:day].to_i

  if day < 1 || day > 25
    puts "Error: Day must be between 1 and 25"
    exit 1
  end

  # Create the day directory
  day_dir = File.expand_path("pkg/day#{sprintf('%02d', day)}", __dir__)
  FileUtils.mkdir_p(day_dir)

  # Copy testdata folder
  FileUtils.cp_r(File.expand_path("template/testdata", __dir__), day_dir)

  # Process ERB templates
  Dir[File.expand_path("template/*.erb", __dir__)].each do |template_path|
    basename = File.basename(template_path, ".erb")
    template = File.read(template_path)
    erb = ERB.new(template, trim_mode: "-")
    # Create a context object with the day variable
    context = OpenStruct.new(day: day)
    output = erb.result(context.instance_eval { binding })

    output_path = File.join(day_dir, basename.sub(/dayXX/, "day#{sprintf('%02d', day)}"))
    File.write(output_path, output)
  end

  puts "Generated files for day #{day} in #{day_dir}"
end
