from selenium import webdriver
from selenium.webdriver import ChromeOptions
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.action_chains import ActionChains as AC
from selenium.webdriver.support import expected_conditions as EC
import time
# import pandas as pd

def new_driver():
    """
    创建新driver对象
    """
    option = webdriver.ChromeOptions()
    option.add_experimental_option('excludeSwitches', ['enable-automation'])
    driver = webdriver.Chrome("./chromedriver",options=option)

    # 设置代理
    # proxy = ""
    # option.add_argument('--proxy-server=%s'%PROXY)
    
    # 隐藏webdriver属性
    driver.execute_cdp_cmd("Page.addScriptToEvaluateOnNewDocument", {
            "source": """Object.defineProperty(navigator, 'webdriver', {
                get: () => undefined
            });"""
        })
    return driver


class fangdiSpdier():
    def __init__(self, url_list, time=3, path="./"):
        self.url_list = url_list # 链接集合
        self.time = time # 默认重试次数
        self.path = "./"

    def request(self,url):
        """
        缓存页面内容
        """
        driver = new_driver()
        driver.get(url)

        # 加载cookies
        cookies = {
            'www.fangdi.com_http_ic':'www.fangdi.com.cn_80_RS', 
            'JSESSIONID1':'NDpliI2IpaaZgyiUezIor0Crhef5gw7_YVZSMz9eAxe_KUy7Y1M-!457065254',
            'www.fangdi.com.cn':'www.fangdi.com.cn_rs1',
        }
        for k in cookies:
            driver.add_cookie({"name":k,"value":cookies[k]})

        
        
        self.retry(driver,3)
        driver.quit()  

    def retry(self,driver,time):
        """
        无内容递归重传
        """
        # 临界值
        if time <= 0 :
            return -1

        # 刷新页面重新加载cookies
        driver.refresh()

        try:
            WebDriverWait(driver,10,0.5).until(
                EC.invisibility_of_element_located((By.CLASS_NAME,"side_bar"))
            )
            self.get_cache(driver)
        except Exception as e:
            print(e)
            self.retry(driver,time-1)

    def do(self):
        for url in self.url_list:
            self.request(url)

    # 缓存当前页面
    def get_cache(self,driver):
        filename = driver.current_url.split("?")[-1]
        with open(self.path + filename + '.html','w') as f:
            f.write(driver.page_source)
        
        with open(self.path + "cookies.txt","w") as f:
            f.writelines(driver.get_cookie())




if __name__ == "__main__":
    url_list = [
        "http://www.fangdi.com.cn/new_house/more_info.html?project_id=c90bb18e9dda0ff9&building_id=54393757b6a130d65f4f10816e5ff473&start_id=04d8b992394ab666",
        "http://www.fangdi.com.cn/new_house/more_info.html?project_id=d0370fc302b06582&building_id=10d7b60946bbdcd28a045eb1690b4806&start_id=14ab23415ba59319",
    ]
    
    spider = fangdiSpdier(url_list)
    spider.do()
    