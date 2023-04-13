
# 支付对接

## 支付宝支付


### 参考文档
- [文档中心](https://b.alipay.com/page/doccenter)
- [在线调试](https://open.alipay.com/api/apiDebug?frontProdCode=I1080300001000041313&backProdCode=I1011000100000000004&apiNames=alipay.trade.app.pay)
- [密钥工具下载](https://opendocs.alipay.com/common/02kipk?pathHash=0d20b438)
- [电脑网站支付](https://opendocs.alipay.com/open/270/105898?pathHash=b3b2b667&ref=api)
- [统一下单接口](https://opendocs.alipay.com/open/028r8t?pathHash=8e24911d&ref=api&scene=22)
- [报错参考](https://opendocs.alipay.com/support/01raxa)
- [接口报错排查](https://openhome.alipay.com/api/errCheck)
- [支付宝支付golang SDK](https://github.com/smartwalle/alipay/)


## 微信支付


### 参考文档
Note:每种支付有自己的适用场景,具体的请自行查阅文档选择
- [文档中心](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pages/index.shtml)
- [Native支付](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml)
- [H5支付接口](https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter2_6_3.shtml)
- [微信证书文档](https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/wechatpay5_0.shtml)
- [微信证书生成证书序列号文档](https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay7_0.shtml) => 生成`openssl x509 -in apiclient_cert.pem -noout -serial`
- [微信支付申请文档](https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter2_5_1.shtml#part-4)
- [Native预支付返回二维码解析网站](https://www.wwei.cn/)
- [回调内容解密](https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/zheng-shu-he-hui-tiao-bao-wen-jie-mi)
