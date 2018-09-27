# Alipay Go Software Development Kit

## 安装

```bash
go get -u github.com/starudream/alipay-sdk-go
```

## 使用

| 已测试接口                                                                                                       |
| --------------------------------------------------------------------------------------------------------------- |
| [publicAppAuthorize.htm](./test/alipay_oauth_public_app_auth_url_test.go)                                       |
| [appToAppAuth.htm](./test/alipay_oauth_app_to_app_auth_url_test.go)                                             |
| [alipay.data.dataservice.bill.downloadurl.query](./test/alipay_data_dataservice_bill_downloadurl_query_test.go) |
| [alipay.fund.auth.operation.cancel](./test/alipay_fund_auth_operation_cancel_test.go)                           |
| [alipay.fund.auth.operation.detail.query](./test/alipay_fund_auth_operation_detail_query_test.go)               |
| [alipay.fund.auth.order.freeze](./test/alipay_fund_auth_order_freeze_test.go)                                   |
| [alipay.fund.auth.order.unfreeze](./test/alipay_fund_auth_order_unfreeze_test.go)                               |
| [alipay.fund.auth.order.voucher.create](./test/alipay_fund_auth_order_voucher_create_test.go)                   |
| [alipay.fund.coupon.operation.query](./test/alipay_fund_coupon_operation_query_test.go)                         |
| [alipay.fund.coupon.order.agreement.pay](./test/alipay_fund_coupon_order_agreement_pay_test.go)                 |
| [alipay.fund.coupon.order.app.pay](./test/alipay_fund_coupon_order_app_pay_test.go)                             |
| [alipay.fund.coupon.order.disburse](./test/alipay_fund_coupon_order_disburse_test.go)                           |
| [alipay.fund.trans.order.query](./test/alipay_fund_trans_order_query_test.go)                                   |
| [alipay.fund.trans.toaccount.transfer](./test/alipay_fund_trans_toaccount_transfer_test.go)                     |
| [alipay.marketing.campaign.cash.create](./test/alipay_marketing_campaign_cash_create_test.go)                   |
| [alipay.marketing.campaign.cash.detail.query](./test/alipay_marketing_campaign_cash_detail_query_test.go)       |
| [alipay.marketing.campaign.cash.list.query](./test/alipay_marketing_campaign_cash_list_query_test.go)           |
| [alipay.marketing.campaign.cash.status.modify](./test/alipay_marketing_campaign_cash_status_modify_test.go)     |
| [alipay.marketing.campaign.cash.trigger](./test/alipay_marketing_campaign_cash_trigger_test.go)                 |
| [alipay.open.auth.token.app](./test/alipay_open_auth_token_app_test.go)                                         |
| [alipay.open.auth.token.app.query](./test/alipay_open_auth_token_app_query_test.go)                             |
| [alipay.pass.instance.add](./test/alipay_pass_instance_add_test.go)                                             |
| [alipay.pass.instance.update](./test/alipay_pass_instance_update_test.go)                                       |
| [alipay.pass.template.add](./test/alipay_pass_template_add_test.go)                                             |
| [alipay.pass.template.update](./test/alipay_pass_template_update_test.go)                                       |
| [alipay.system.oauth.token](./test/alipay_system_oauth_token_test.go)                                           |
| [alipay.trade.app.pay](./test/alipay_trade_app_pay_test.go)                                                     |
| [alipay.trade.cancel](./test/alipay_trade_cancel_test.go)                                                       |
| [alipay.trade.close](./test/alipay_trade_close_test.go)                                                         |
| [alipay.trade.create](./test/alipay_trade_create_test.go)                                                       |
| [alipay.trade.fastpay.refund.query](./test/alipay_trade_fastpay_refund_query_test.go)                           |
| [alipay.trade.page.pay](./test/alipay_trade_page_pay_test.go)                                                   |
| [alipay.trade.pay](./test/alipay_trade_pay_test.go)                                                             |
| [alipay.trade.precreate](./test/alipay_trade_precreate_test.go)                                                 |
| [alipay.trade.query](./test/alipay_trade_query_test.go)                                                         |
| [alipay.trade.refund](./test/alipay_trade_refund_test.go)                                                       |
| [alipay.trade.wap.pay](./test/alipay_trade_wap_pay_test.go)                                                     |
| [alipay.user.info.auth](./test/alipay_user_info_auth_test.go)                                                   |
| [alipay.user.info.share](./test/alipay_user_info_share_test.go)                                                 |
| [monitor.heartbeat.syn](./test/monitor_heartbeat_syn_test.go)                                                   |

## 作者

-   [StarUDream](https://github.com/starudream)

## 开源协议

-   [MIT](./LICENSE)

## 其他资料

-   [蚂蚁金服开放平台 API 文档](https://docs.open.alipay.com/api)
