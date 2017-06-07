import os
import sys
import new
import unittest
import time
from selenium import webdriver
from sauceclient import SauceClient
# it's best to remove the hardcoded defaults and always get these values
# from environment variables
USERNAME = os.environ['SAUCE_USER']
ACCESS_KEY = os.environ['SAUCE_KEY']
sauce = SauceClient(USERNAME, ACCESS_KEY)
browsers = [{"platform": "Mac OS X 10.9",
             "browserName": "chrome",
             "version": "58"},
            {"platform": "Windows 10",
             "browserName": "firefox",
             "version": "53"}]

def on_platforms(platforms):
    def decorator(base_class):
        module = sys.modules[base_class.__module__].__dict__
        for i, platform in enumerate(platforms):
            d = dict(base_class.__dict__)
            d['desired_capabilities'] = platform
            name = "%s_%s" % (base_class.__name__, i + 1)
            module[name] = new.classobj(name, (base_class,), d)
    return decorator

@on_platforms(browsers)
class SauceSampleTest(unittest.TestCase):
    def setUp(self):
        self.desired_capabilities['name'] = self.id()
        sauce_url = "http://%s:%s@ondemand.saucelabs.com:80/wd/hub"
        self.driver = webdriver.Remote(
            desired_capabilities=self.desired_capabilities,
            command_executor=sauce_url % (USERNAME, ACCESS_KEY)
        )
        self.driver.implicitly_wait(30)
    def test_sauce(self):
        self.driver.get(os.environ['HEROKU_URL'])
        assert "Random Data Grapher" in self.driver.title

        # Wait until dot value is above 0
        time.sleep(5)
        dotCount = int(self.driver.find_element_by_id('dotCount').text)
        assert dotCount >= 40
    def tearDown(self):
        print("Link to your job: https://saucelabs.com/jobs/%s" % self.driver.session_id)
        try:
            if sys.exc_info() == (None, None, None):
                sauce.jobs.update_job(self.driver.session_id, passed=True)
            else:
                sauce.jobs.update_job(self.driver.session_id, passed=False)
        finally:
            self.driver.quit()

if __name__ == '__main__':
    unittest.main()
