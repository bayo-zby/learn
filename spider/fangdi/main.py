from selenium import webdriver
from selenium.webdriver import ChromeOptions
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.action_chains import ActionChains as AC
from selenium.webdriver.support import expected_conditions as EC
import pandas as pd
import os,json
from datetime import datetime
# import random

def new_driver():
    """
    创建新driver对象
    """
    option = webdriver.ChromeOptions()
    option.add_experimental_option('excludeSwitches', ['enable-automation']) # 取消信息栏
    # option.add_argument('--headless')     # 无头模式
    # option.add_argument('--start-maximized') # 最大化
    # option.add_argument(f'--proxy-server={self.proxy}') # 设置代理
    driver = webdriver.Chrome("./chromedriver",options=option) 

    # 隐藏webdriver属性,跳过检测
    driver.execute_cdp_cmd("Page.addScriptToEvaluateOnNewDocument", {
            "source": """Object.defineProperty(navigator, 'webdriver', {
                get: () => undefined
            });"""
        })
    return driver


class fangdiSpdier():
    """
    @param  url_list    访问链接集合
    @param  time        重试次数
    @path   文件路径，包括缓存文件夹和cookie
    """
    def __init__(self, url_list, time=3, path="./"):
        self.driver = None
        self.url_list = url_list # 链接集合
        self.time = time # 重试次数,默认为三次
        self.path = path
        self.cache_dir = self.path + "/cache/"
        
        if os.path.exists(self.cache_dir) == False:
            os.mkdir(self.cache_dir)
        # IP池
        # self.proxy = random.choice([
        # ])

    def writer_log(self,level,content):
        """
        @ brief : 保持log文件
        @ param : level,日志级别
        @ content : content，日志主题内容
        """
        time = datetime.now() # 本地时间
        with open("log.txt","a") as f:
            f.write(f"{time} {level}:{content}\n")

    def request(self,url,filepath,retry):
        """
        @ brief : 访问并缓存页面内容,由于在未加载出数据的页面刷新有时候
                  会导致空白页，因此选择重置模拟浏览器
        @ param ： url,访问链接
        @ param ： filepath,文件储存路径
        """
        self.writer_log("TRACK",f"访问链接：{url}")
        while self.time >= retry:
            self.writer_log("TRACK",f"第{retry}次尝试")
            self.driver = new_driver()
            self.driver.get(url)
            cookie = self.set_cookies() # 载入cookie
            if self.webwait(): 
                self.writer_log("TRACK",f"访问成功，开始进行页面缓存")
                self.get_cache(filepath,cookie)
                break
            retry += 1
        if retry > self.time:
            self.writer_log("ERROR",f"超过重试次数，跳过当前链接")

        self.driver.quit()
        

    def set_cookies(self):
        """
        @ brief : cookie加载，完毕后进行刷新
        @ ret ： 返回cookie，做刷新
        """
        self.driver.add_cookie({"name":"www.fangdi.com.cn","value":"www.fangdi.com.cn_rs1"})
        self.driver.add_cookie({"name":"www.fangdi.com_http_ic","value":"www.fangdi.com.cn_80_RS"})

        with open(self.path + "cookies.txt","r") as f:
            cookie = json.load(f)
            # 加载cookies
            self.writer_log("TRACK",f"加载COOKIE {cookie}")
            self.driver.add_cookie(cookie)
            
        # 刷新页面加载cookie
        self.driver.refresh()

        return cookie


    def webwait(self):
        """
        @ brief : 等待匹配内容，访问失败，则刷新重试
        @ ret   : 返回0或1来表示页面访问是否成功
        """
        try:
            # 判断页面加载完成
            WebDriverWait(self.driver,20,1).until(
                lambda driver:self.driver.find_element_by_xpath("//div[@id='moreInfo']/div[@class='lay1']/table")
            )
            return 1
        except:
            self.writer_log("ERROR",f"未发现匹配内容或加载超时")
            return 0
            
        

    def do(self):
        """
        依次运行列表内容
        已经有缓存的链接跳过
        """

        # 后期加入多线程操作
        for url in self.url_list:
            try:
                filename = url.split("?")[-1]
                filepath = f'{self.cache_dir}/{self.path + filename}.html'
                # 链接去重
                if os.path.exists(filepath):
                    self.writer_log("TRACK",f"已抓取，跳过：{url}")
                    continue
                
                # 同一浏览器重复访问会导致主页跳转至主页
                self.request(url,filepath,1)
                
            except Exception as e:
                self.writer_log("ERROR",f"抓取程序出现未知错误：{e}")


    
    def get_cache(self,filepath,cookie_ori):
        """
        @ brief : 缓存当前页面
        @ param : filepath,缓存文件地址,self.cache/链接所有参数.html    
        """
        # 缓存页面
        with open(filepath,'w') as f:
            f.write(self.driver.page_source)

        # 缓存新的cookie
        with open(self.path + "cookies.txt","w") as f:
            cookies = self.driver.get_cookies()
            for cookie in  cookies:
                if cookie["name"] == "JSESSIONID1" and cookie["value"] != cookie_ori["value"]:
                    f.write(json.dumps(cookie))
        self.writer_log("TRACK",f"内容缓存完毕")

    
    
def getUrl(filepath):
    """
    @brief  访问xlsx文件，然后读取第二列内容
    @param  filepath,文件路径
    @ret    list,链接列表
    """
    data = pd.read_excel(filepath,header=None)[1][1:]
    return [i for i in data]

if __name__ == "__main__":
    # url = 'http://www.fangdi.com.cn/'

    url_list = getUrl('url.xlsx')
    # url_list = ['http://www.fangdi.com.cn/new_house/new_house_list.html?districtID=27d3af3bd45acf5e&houseAreaID=0&houseTypeID=0&dicAvgpriceID=0']
    spider = fangdiSpdier(url_list)
    spider.do()

