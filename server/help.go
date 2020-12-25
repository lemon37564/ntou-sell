package server

// API is the backend APIs
const API = `
<html>
	<H1>後端API</H1>
	<H4>
		<body>
    		<form action="/backend/user/login" method="post">
      			Account:<input type="text" name="account">
      			Password:<input type="password" name="password">
      			<input type="submit" value="Submit">
    		</form>
  		</body>
	<a href=/backend/user/logout>登出</a>
	<br>
	</H4>
	
	<a href=/backend/user/help>關於使用者的功能...</a><br><br>

	<a href=/backend/product/help>關於商品的功能...</a><br><br>

	<a href=/backend/history/help>關於歷史紀錄的功能...</a><br><br>
	
	<a href=/backend/order/help>關於訂單的功能...</a><br><br>

	<a href=/backend/bid/help>關於競標的功能...</a><br><br>

	<a href=/backend/cart/help>關於購物車的功能...</a><br><br>
	
	關於販賣商品的功能...(此功能已整合到product)<br><br>
	
	<a href=/backend/message/help>關於使用者對話的功能...</a><br><br>

	<a href=/backend/pics/help>關於圖片的功能...</a><br><br>
</html>
`

// UserAPI is API of user functions
const UserAPI = `
<html>
	<h3>使用者的所有功能皆使用POST傳輸</h3>
	<p>
		/backend/user/all<br>
		列出所有帳號(此功能已關閉，請註冊新帳號)<br><br>

		/backend/user/login<br>
		登入是否成功<br>
		參數: "account", "password"
		<br><br>

		/backend/user/logout<br>
		登出功能<br>
		<a href=/backend/user/logout>
		/backend/user/logout</a>
		<br><br>

		/backend/user/regist<br>
		註冊新帳號<br>
		參數: "account", "password", "name"<br><br>

		/backend/user/delete<br>
		刪除帳號<br>
		參數: "password"
		<br><br>

		/backend/user/delete<br>
		更換目前登入帳號的密碼<br>
		參數: "oldPassword", "newPassword"
		<br><br>

		/backend/user/delete<br>
		更換目前登入帳號的使用者姓名(暱稱)<br>
		參數: "newName"
		<br><br>
	</p>
</html>
`

// ProductAPI is API of product functions
const ProductAPI = `
<html>
	<p> 
		/backend/product/urproduct<br>
		列出自己販賣中的商品<br>
		<br><br>

		/backend/product/newest?amount=...<br>
		e.g.顯示最新商品(3筆資料)<br>
		<a href=/backend/product/newest?amount=3>
		/backend/product/newest?amount=3</a>
		<br><br>

		暫時新增的上傳商品(POST)
		/backend/product/postadd<br>
		新增商品以及圖片<br>
		參數:<br>
		name<br>
		price<br>
		description<br>
		amount<br>
		bid<br>
		date<br>
		uploadefile(圖片)<br>
		參數內容可以參考一般的add
		<br><br>

		/backend/product/add?name=...&price=...&description=...&amount=...&bid=...&date=...<br>
		新增商品<br>
		e.g.新增一商品->商品名:ifone12價格:5000，商品說明:盜版商品，競標:是，競標期限:2006-01-02<br>
		<a href=/backend/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&bid=true&date=2006-01-02>
		/backend/product/add?name=ifone12&price=5000&description=盜版商品&amount=10&bid=true&date=2006-01-02</a>
		<br><br>

		/backend/product/get?pdid=...<br>
		取得商品資訊(使用商品id取)<br>
		e.g.取得商品id為1的資訊<br>
		<a href=/backend/product/get?pdid=1>
		/backend/product/get?pdid=1</a>
		<br><br>

		/backend/product/search?name=...<br>
		查詢商品(一般)<br>
		e.g.查詢商品名中含有"ifone"的商品<br>
		<a href=/backend/product/search?name=ifone>
		/backend/product/search?name=ifone</a>
		<br>
		查詢商品(過濾)(參數任選)<br>
		/backend/product/search?name=...&minprice=...&maxprice=...&eval=...<br>
		e.g.查詢商品名中含有"ifone"的商品，最低價格為10，最高價格為5000000，最低評價為0<br>
		<a href=/backend/product/search?name=ifone&minprice=10&maxprice=5000000&eval=0>
		/backend/product/search?name=ifone&minprice=10&maxprice=5000000&eval=0</a>
		<br><br>
	</p>
</html>
`

// HistoryAPI is API of history funcions
const HistoryAPI = `
<html>
	<p> 
		/backend/history/all<br>
		列出歷史紀錄(此功能已關閉)<br><br>

		/backend/history/add?pdid=...<br>
		增加一筆新的歷史紀錄<br>
		e.g.新增目前登入帳號之商品id為1的歷史紀錄<br>
		<a href=/backend/history/add?pdid=1>
		/backend/history/add?pdid=1</a>
		<br><br>
	
		/backend/history/get?amount=...&newest=...<br>
		查詢歷史紀錄(newest=true則最新的紀錄在前面，newest=false則反之)<br>
		e.g.查詢目前登入帳號的歷史紀錄<br>
		<a href=/backend/history/get?amount=10&newest=true>
		/backend/history/get?amount=10&newest=true</a>
		<br><br>

		/backend/history/delete?pdid=...<br>
		刪除歷史紀錄<br>
		e.g.刪除目前登入的帳號中商品編號為2的歷史紀錄<br>
		<a href=/backend/history/delete?pdid=2>
		/backend/history/delete?pdid=2</a>
		<br><br>

		/backend/history/delete?pdids=...,...,...,...<br>
		刪除特定歷史紀錄(複數)<br>
		e.g.刪除目前登入的帳號中商品編號為1,2,3,4的歷史紀錄<br>
		<a href=/backend/history/delete?pdids=1,2,3,4>
		/backend/history/delete?pdids=1,2,3,4</a>
		<br><br>

		/backend/history/deleteall<br>
		刪除所有歷史紀錄<br>
		e.g.刪除目前登入的帳號的所有歷史紀錄<br>
		<a href=/backend/history/deleteall>
		/backend/history/deleteall</a>
		<br><br>
	</p>
</html>
`

// OrderAPI is API of order functions
const OrderAPI = `
<html>
	<p>
		/backend/order/get<br>
		取得使用者訂單資訊<br>
		e.g.使用者進入訂單時顯示他購買東西 (使用前要先買喔)<br>
		<a href=/backend/order/get>
		/backend/order/get</a>
		<br><br>

		/backend/order/add?pdid=...&amount=...<br>
		把商品加入訂單<br>
		e.g.使用者在cart點選購買時可以加入訂單<br>
		<a href=/backend/order/add?pdid=2&amount=2>
		/backend/order/add?pdid=2&amount=2</a>
		<br><br>

		/backend/order/del?pdid=...<br>
		把商品從訂單中刪除<br>
		e.g.使用者可以把order裡的東西刪掉(這個需要改，應該要買賣家溝通才能刪)<br>
		<a href=/backend/order/del?pdid=2>
		/backend/order/del?pdid=2</a>
		<br><br>
	</p>
</html>
`

// BidAPI is API of bid functions
const BidAPI = `
<html>
	<p>
		/backend/bid/get?pdid=...<br>
		在商品頁面取得競標商品資訊<br>
		e.g.商品頁面選取競標資訊<br>
		<a href=/backend/bid/get?pdid=6>
		/backend/bid/get?pdid=6</a>
		<br><br>

		/backend/bid/set?pdid=...&money=...<br>
		更新競標資訊<br>
		e.g.買家競標了商品<br>
		<a href=/backend/bid/set?pdid=6&money=1000>
		/backend/bid/set?pdid=6&money=1000</a>
		<br><br>

		/backend/bid/delete?pdid=...<br>
		刪除競標商品<br>
		e.g.在競標商品已被購買情況下，前端呼叫此功能<br>
		<a href=/backend/bid/delete?pdid=6>
		/backend/bid/delete?pdid=6</a>
		<br><br>
	</p>
</html>
`

// CartAPI is API of cart functions
const CartAPI = `
<html>
	<p>
		/backend/cart/add?pdid=...&amount=...<br>
		加入購物車<br>
		e.g.商品頁面確定購買時須用到此功能<br>
		<a href=/backend/cart/add?pdid=2&amount=3>
		/backend/cart/add?pdid=2&amount=3</a>
		<br><br>

		/backend/cart/remo?pdid=...<br>
		刪除在購物車的商品<br>
		e.g.<br>
		<a href=/backend/cart/remo?pdid=2>
		/backend/cart/remo?pdid=2</a>
		<br><br>

		/backend/cart/modf?pdid=...&amount=...<br>
		在購物車更改數量<br>
		e.g.<br>
		<a href=/backend/cart/modf?pdid=2&amount=4>
		/backend/cart/modf?pdid=2&amount=4</a>
		<br><br>
	
		/backend/cart/tal<br>
		回傳目前在購物車選取物品的總金額<br>
		e.g.<br>
		<a href=/backend/cart/tal>
		/backend/cart/tal</a>
		<br><br>

		/backend/cart/geps<br>
		回傳放在購物車的商品們<br>
		e.g.<br>
		<a href=/backend/cart/geps>
		/backend/cart/geps</a>
		<br><br>
	</p>
</html>
`

// MessageAPI is API of message functions
const MessageAPI = `
<html>
	<p>
		/backend/message/all<br>
		所有聊天紀錄(僅限開發期間)<br>
		<a href=/backend/message/all>
		/backend/message/all</a>
		<br><br>

		/backend/message/send?remoteUID=...&text=...<br>
		新增聊天紀錄<br>
		e.g.新增聊天紀錄<br>
		<a href=/backend/message/send?remoteUID=2&text=你好>
		/backend/message/send?remote=2&text=你好</a>
		<br><br>

		/backend/message/get?remoteUID=...&ascend=...<br>
		取得聊天紀錄，(ascend=true會讓最新的在最前面，ascend=false則反之)<br>
		內有參數Status，若為s表示send，r表示receive<br>
		e.g.取得跟某用戶的聊天紀錄(有照時間順序排好)<br>
		<a href=/backend/message/get?remoteUID=2&ascend=true>
		/backend/message/get?remoteUID=2&ascend=true</a>
		<br><br>
	</p>
</html>
`

// PicAPI is API of picserver and pictrue functions
const PicAPI = `
<html>
	<p>
		/img/<br>
		列出所有圖片(僅限開發期間)<br>
		<a href=/img/>
		/img/</a><br>
		查看特定圖片時，網址為/img/檔名<br>
		<a href=/img/server.jpg>範例:/img/server.jpg</a>
		<br><br>

		/backend/pics/get?pdid=...<br>
		查詢圖片名稱<br>
		回傳圖片路徑<br>
		<br><br>
	
		/backend/pics/upload<br>
		上傳圖片(用post)<br>

		前端範例:<br>
		<form enctype="multipart/form-data" action="/backend/pics/upload" method="post"><br>
		<input type="file" name="uploadfile" /><br>
		<input type="hidden" name="token" value="{{.}}"/><br>
		<input type="submit" value="upload" /><br>
		<br><br>
	</p>
</html>
`
