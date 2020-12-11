package server

const helpPage = `
<html>
	<H1>後端API</H1>
	<H4>測試用帳密:account=1234&password=1234<br>
	<a href=https://se-ssb.herokuapp.com/user/login?account=1234&password=1234>登入</a><br>
	</H4>
	<p>
		/user/all<br>
		列出所有帳號(僅限開發期間)<br>
		<a href=/user/all> /user/all </a><br><br>
	</p>
	<p> 
		/history/add?account=...&password=...<br>
		登入是否成功(bool)<br>
		e.g.登入帳號為test@gmail.com以及密碼為0000的使用者<br>
		<a href=https://se-ssb.herokuapp.com/user/login?account=test@gmail.com&password=0000>
		https://se-ssb.herokuapp.com/user/login?account=test@gmail.com&password=0000</a>
		<br><br>
	</p>
	<p>
		/user/regist?account=...&password=...&name=...<br>
		註冊新帳號<br>
		e.g.註冊一帳號為test2@gmail.com，密碼為1234，使用者姓名為Wilson的帳號<br>
		<a href=https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson>
		https://se-ssb.herokuapp.com/user/regist?account=test2@gmail.com&password=1234&name=Wilson</a>
		<br><br>
	</p>
	<p>
		/user/delete?account=...&password=...<br>
		刪除帳號<br>
		e.g.刪除帳號為test3@gmail.com的帳號(需要輸入密碼驗證:密碼為1234)<br>
		<a href=https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234>
		https://se-ssb.herokuapp.com/user/delete?account=test3@gmail.com&password=1234</a>
		<br><br>
	</p>
	<p> 
		/product/all<br>
		列出所有商品(僅限開發期間)<br>
		<a href=/product/all> /product/all </a><br><br>
	</p>
	<p>
		/product/newest?amount=...<br>
		e.g.顯示最新商品(3筆資料)<br>
		<a href=https://se-ssb.herokuapp.com/product/newest?amount=3>
		https://se-ssb.herokuapp.com/product/newest?amount=3</a>
		<br><br>
	<p> 
		/product/add?name=...&price=...&description=...&amount=...&account=...&bid=...&date=...<br>
		新增商品<br>
		e.g.新增一商品->商品名:ifone12價格:5000，商品說明:盜版商品，帳號:test@gmail.com，競標:是，競標期限:2020-12-31<br>
		<a href=https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31>
		https://se-ssb.herokuapp.com/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&account=test@gmail.com&bid=true&date=2020-12-31</a>
		<br><br>
	</p>
	<p>
		/product/search?name=...<br>
		查詢商品<br>
		e.g.查詢商品名中含有"ifone"的商品<br>
		<a href=https://se-ssb.herokuapp.com/product/search?name=ifone>
		https://se-ssb.herokuapp.com/product/search?name=ifone</a>
		<br><br>
	</p>
	<p>
		/product/filterSearch?name=...&minprice=...&maxprice=...&eval=...<br>
		查詢商品(過濾)<br>
		e.g.查詢商品名中含有"ifone"的商品，最低價格為10，最高價格為5000，最低評價為2<br>
		<a href=https://se-ssb.herokuapp.com/product/filterSearch?name=ifone&minprice=10&maxprice=5000&eval=2>
		https://se-ssb.herokuapp.com/product/filterSearch?name=ifone&minprice=10&maxprice=5000&eval=2</a>
		<br><br>
	</p>
	<p> 
		/history/all<br>
		列出歷史紀錄(僅限開發期間)<br>
		<a href=/history/all> /history/all </a><br><br>
	</p>
	<p> 
		/history/add?account=...&pdid=...<br>
		增加一筆新的歷史紀錄<br>
		e.g.新增帳號為test@gmail.com以及商品id為1的歷史紀錄<br>
		<a href=https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1>
		https://se-ssb.herokuapp.com/history/add?account=test@gmail.com&pdid=1</a>
		<br><br>
	</p>
	<p>
		/history/get?account=...&amount=...<br>
		查詢歷史紀錄<br>
		e.g.查詢帳號為test2@gmail.com的10歷史紀錄<br>
		<a href=https://se-ssb.herokuapp.com/history/get?account=test2@gmail.com&amount=10>
		https://se-ssb.herokuapp.com/history/get?account=test2@gmail.com&amount=10</a>
		<br><br>
	</p>
	<p>
		/history/delete?account=...&pdid=...<br>
		刪除歷史紀錄<br>
		e.g.刪除帳號test3@gmail.com以及商品編號為2的歷史紀錄<br>
		<a href=https://se-ssb.herokuapp.com/history/delete?account=test3@gmail.com&pdid=2>
		https://se-ssb.herokuapp.com/history/delete?account=test3@gmail.com&pdid=2</a>
		<br><br>
	</p>
	<p>
		/order/get?uid=...<br>
		取得使用者訂單資訊<br>
		e.g.使用者進入訂單時顯示他購買東西 (使用前要先買喔)<br>
		<a href=https://se-ssb.herokuapp.com/order/get?uid=1>
		https://se-ssb.herokuapp.com/history/order/get?uid=1</a>
		<br><br>
	</p>
	<p>
		/order/add?uid=...&pdid=...&amount=...<br>
		把商品加入訂單<br>
		e.g.使用者在cart點選購買時可以加入訂單<br>
		<a href=https://se-ssb.herokuapp.com/order/add?uid=1&pdid=2&amount=2>
		https://se-ssb.herokuapp.com/order/add?uid=1&pdid=2&amount=2</a>
		<br><br>
	</p>
	<p>
		/order/del?uid=...&pdid=...<br>
		把商品從訂單中刪除<br>
		e.g.使用者可以把order裡的東西刪掉(這個需要改，應該要買賣家溝通才能刪)<br>
		<a href=https://se-ssb.herokuapp.com/order/del?uid=1&pdid=2>
		https://se-ssb.herokuapp.com/order/del?uid=1&pdid=2</a>
		<br><br>
	</p>
	<p>
		/bid/get?pdid=...<br>
		在商品頁面取得競標商品資訊<br>
		e.g.商品頁面選取競標資訊<br>
		<a href=https://se-ssb.herokuapp.com/bid/get?pdid=6>
		https://se-ssb.herokuapp.com/bid/get?pdid=6</a>
		<br><br>
	</p>
	<p>
		/bid/set?pdid=...&uid=...&money=...<br>
		更新競標資訊<br>
		e.g.買家競標了商品<br>
		<a href=https://se-ssb.herokuapp.com/bid/set?pdid=6&uid=1&money=1000>
		https://se-ssb.herokuapp.com/bid/set?pdid=6&uid=1&money=1000</a>
		<br><br>
	</p>
	<p>
		/bid/delete?pdid=...<br>
		刪除競標商品<br>
		e.g.在競標商品已被購買情況下，前端呼叫此功能<br>
		<a href=https://se-ssb.herokuapp.com/bid/delete?pdid=6>
		https://se-ssb.herokuapp.com/bid/delete?pdid=6</a>
		<br><br>
	</p>
	<p>
		/cart/add?uid=...&pdid=...&amount=...<br>
		加入購物車<br>
		e.g.商品頁面確定購買時須用到此功能<br>
		<a href=https://se-ssb.herokuapp.com/cart/add?uid=1&pdid=2&amount=3>
		https://se-ssb.herokuapp.com/cart/add?uid=1&pdid=2&amount=3</a>
		<br><br>
	</p>
	<p>
		/cart/remo?uid=...&pdid=...<br>
		刪除在購物車的商品<br>
		e.g.<br>
		<a href=https://se-ssb.herokuapp.com/cart/remo?uid=1&pdid=2>
		https://se-ssb.herokuapp.com/cart/remo?uid=1&pdid=2</a>
		<br><br>
	</p>
	<p>
		/cart/modf?uid=...&pdid=...&amount=...<br>
		在購物車更改數量<br>
		e.g.<br>
		<a href=https://se-ssb.herokuapp.com/cart/modf?uid=1&pdid=2&amount=4>
		https://se-ssb.herokuapp.com/cart/modf?uid=1&pdid=2&amount=4</a>
		<br><br>
	</p>
	<p>
		/cart/tal?uid=...<br>
		回傳目前在購物車選取物品的總金額<br>
		e.g.<br>
		<a href=https://se-ssb.herokuapp.com/cart/tal?uid=1>
		https://se-ssb.herokuapp.com/cart/tal?uid=1</a>
		<br><br>
	</p>
	<p>
		/cart/geps?uid=...<br>
		回傳放在購物車的商品們<br>
		e.g.<br>
		<a href=https://se-ssb.herokuapp.com/cart/geps?uid=1>
		https://se-ssb.herokuapp.com/cart/geps?uid=1</a>
		<br><br>
	</p>
	<p>
		/sell/set?pdname=...&price=...&description=...&amount=...&account=...&sellerID=...&bid=...&date=...&dateLine=...<br>
		販賣商品<br>
		e.g.在販賣網頁販賣商品<br>
		<a href=https://se-ssb.herokuapp.com/sell/set?pdname="火箭"&price="50"&description="這是火箭"&amount="3"&account="test@gmail.com"&sellerID="2"&bid="true"&date="1229"&dateLine="1231">
		https://se-ssb.herokuapp.com/sell/set?pdname="火箭"&price="50"&description="這是火箭"&amount="3"&account="test@gmail.com"&sellerID="2"&bid="true"&date="1229"&dateLine="1231"</a>
		<br><br>
	</p>
</html>
`
