require 'minitest/autorun'
require_relative 'niceness'

class NicenessTest < Minitest::Test
  def test_ugknbfddgicrmopn_nice_v1
    assert 'ugknbfddgicrmopn'.nice?(Niceness::V1)
  end

  def test_aaa_nice_v1
    assert 'aaa'.nice?(Niceness::V1)
  end

  def test_jchzalrnumimnmhp_naughty_v1
    assert 'jchzalrnumimnmhp'.naughty?(Niceness::V1)
  end

  def test_haegwjzuvuyypxyu_naughty_v1
    assert 'haegwjzuvuyypxyu'.naughty?(Niceness::V1)
  end

  def test_dvszwmarrgswjxmb_naughty_v1
    assert 'dvszwmarrgswjxmb'.naughty?(Niceness::V1)
  end

  def test_qjhvhtzxzqqjkmpb_nice_v2
    assert 'qjhvhtzxzqqjkmpb'.nice?
  end

  def test_xxyxx_nice_v2
    assert 'xxyxx'.nice?
  end

  def test_uurcxstgmygtbstg_naughty_v2
    assert 'uurcxstgmygtbstg'.naughty?
  end

  def test_ieodomkazucvgmuy_naughty_v2
    assert 'ieodomkazucvgmuy'.naughty?
  end
end
