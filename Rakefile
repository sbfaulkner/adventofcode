require "rake/testtask"

task default: "test"

Rake::TestTask.new do |t|
  t.libs << "pkg"
  t.test_files = FileList["pkg/*/day*_test.rb"]
end
