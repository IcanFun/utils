package i18n

import (
	"redis-trade-query-server/utils"

	"github.com/IcanFun/go-i18n-struct/i18nfile"
)

func init() {
	i18ns := new(utils.I18ns)

	i18ns.AddI18n(utils.I18n{
		Id:    "api.api.init.parsing_templates.debug",
		Zh_CN: "解析服务模板 %v",
		En_US: "Parsing Service Template %v",
		Zh_HK: "解析服務範本 %v",
		Vi_VN: "Dịch vụ phân tích mẫu %v",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.api.init.parsing_templates.error",
		Zh_CN: "解析服务模板出错 %v",
		En_US: "Error parsing service template %v",
		Zh_HK: "解析服務範本出錯 %v",
		Vi_VN: "Dịch vụ phân tích mẫu sai. %v",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "single-sign-on.config_file",
		Zh_CN: "从 %v 加载配置文件",
		En_US: "Loading configuration files from %v",
		Zh_HK: "從 %v 加載設定檔",
		Vi_VN: "Từ %v tải tập tin cấu hình",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "single-sign-on.current_version",
		Zh_CN: "当前版本是 %v (%v/%v/%v)",
		En_US: "The current version is %v (%v/%v/%v)",
		Zh_HK: "當前版本是 %v (%v/%v/%v)",
		Vi_VN: "Phiên bản hiện tại là %v (%v/%v/%v)",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "single-sign-on.working_dir",
		Zh_CN: "当前工作目录是 %v",
		En_US: "The current working directory is %v",
		Zh_HK: "當前工作目錄是 %v",
		Vi_VN: "Thư mục làm việc hiện tại là %v",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    PARAM_ERROR,
		Zh_CN: "参数错误",
		En_US: "Parameter error",
		Zh_HK: "參數錯誤",
		Vi_VN: "Tham số sai",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.context.permissions.app_error",
		Zh_CN: "您没有对应的权限",
		En_US: "Operation not permitted",
		Zh_HK: "您沒有對應的許可權",
		Ja_JP: "対応する許可権がありません",
		Vi_VN: "Bạn không có quyền phù hợp",
		Ru_RU: "У вас нет соответствующих прав",
		Pt_BR: "Você não tem as permissões correspondentes",
		Th_TH: "คุณไม่ได้รับสิทธิ์ที่เกี่ยวข้อง",
		Ko_KR: "해당 권한을 가지고 있지 않습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.server.new_server.init.info",
		Zh_CN: "服务正在初始化...",
		En_US: "The service is being initialized...",
		Zh_HK: "服務正在初始化...",
		Vi_VN: "Dịch vụ đang bắt đầu...",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.server.start_server.listening.info",
		Zh_CN: "服务正在监听 %v",
		En_US: "Services are being monitored %v",
		Zh_HK: "服務正在監聽 %v",
		Vi_VN: "Dịch vụ đang giám sát %v",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.server.start_server.starting.info",
		Zh_CN: "启动服务...",
		En_US: "Startup service...",
		Zh_HK: "啟動服務...",
		Vi_VN: "Dịch vụ khởi động...",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.verify_subject",
		Zh_CN: "[{{ .SiteName }}] 电子邮件验证",
		En_US: "[{{ .SiteName }}] E-mail validation",
		Zh_HK: "[{{ .SiteName }}] 電子郵件驗證",
		Ja_JP: "[{{ .SiteName }}] 電子メール認証",
		Vi_VN: "[{{ .SiteName }}] Xác nhận email",
		Ru_RU: "[{{ .SiteName }}] Подтверждение адреса эл. почты",
		Pt_BR: "[{{ .SiteName }}] Verificação de e-mail",
		Th_TH: "[{{ .SiteName }}] ยืนยันอีเมล",
		Ko_KR: "[{{ .SiteName }}] 이메일 확인",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.message_subject",
		Zh_CN: "[{{ .SiteName }}] 新消息",
		En_US: "[{{ .SiteName }}] New Message",
		Zh_HK: "[{{ .SiteName }}] 新消息",
		Ja_JP: "[{{ .SiteName }}] 新情報",
		Vi_VN: "[{{ .SiteName }}] Tin nhắn mới",
		Ru_RU: "[{{ .SiteName }}] Новое сообщение",
		Pt_BR: "[{{ .SiteName }}] Nova mensagem",
		Th_TH: "[{{ .SiteName }}] ข้อความใหม่",
		Ko_KR: "[{{ .SiteName }}] 새로운 메시지",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.email_verified_code_body.title",
		Zh_CN: "欢迎加入 {{ .SiteName }}",
		En_US: "Welcome to {{ .SiteName }}",
		Zh_HK: "歡迎加入 {{ .SiteName }}",
		Ja_JP: "ようこそ {{ .SiteName }}",
		Vi_VN: "Chào mừng {{ .SiteName }}",
		Ru_RU: "Добро пожаловать {{ .SiteName }}",
		Pt_BR: "Bem vindo {{ .SiteName }}",
		Th_TH: "ยินดีต้อนรับ {{ .SiteName }}",
		Ko_KR: "환영합니다. {{ .SiteName }}",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.email_verified_code_body.info",
		Zh_CN: "请在邮箱注册页面上输入下面的验证码，记住不要轻易把验证码告诉别人",
		En_US: "Please enter the verification code below on the email registration page, remember not to tell the verification code easily.",
		Zh_HK: "請在郵箱註冊頁面上輸入下面的驗證碼，記住不要輕易把驗證碼告訴別人",
		Ja_JP: "メールボックス登録ページで、下の認証コードを入力してください。認証コードは他人に簡単に教えないようにしてください",
		Vi_VN: "Vui lòng nhập mã xác minh bên dưới vào trang đăng ký email, lưu ý không tiết lộ mã xác minh này.",
		Ru_RU: "Введите проверочный код ниже на странице регистрации эл. почты, помните о необходимости держать его в секрете.",
		Pt_BR: "Inserir o código de verificação abaixo na página de registro do e-mail, lembre-se de não informar um código de verificação muito fácil.",
		Th_TH: "กรุณาใส่รหัสยืนยันด้านล่างในหน้าลงทะเบียนอีเมล พึงระลึกว่าไม่ควรบอกรหัสยืนยันให้แก่ใครโดยง่าย",
		Ko_KR: "아래 이메일 등록 페이지에 인증 코드를 입력하세요. 인증 코드를 아무에게나 공유해서는 안 됩니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.email_info",
		Zh_CN: "有任何问题请随时给我们发送电子邮件: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.{{.SiteName}}团队<br>",
		En_US: "Please feel free to email us if you have any questions.: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.{{.SiteName}}Team<br>",
		Zh_HK: "有任何問題請隨時給我們發送電子郵件: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.{{.SiteName}}團隊<br>",
		Ja_JP: "ご質問等がございましたら、いつでもメールにてお問い合わせください: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.{{.SiteName}}チーム<br>",
		Vi_VN: "Vui lòng gửi email cho chúng tôi nếu bạn có bất kỳ câu hỏi nào: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.Đội ngũ {{.SiteName}}<br>",
		Ru_RU: "Не стесняйтесь задавать нам по эл. почте любые интересующие вас вопросы: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.Команда {{.SiteName}}<br>",
		Pt_BR: "Sinta-se à vontade para nos enviar um e-mail se tiver qualquer dúvida: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.Equipe {{.SiteName}}<br>",
		Th_TH: "หากคุณมีข้อสงสัยใด ๆ โปรดอย่าลังเลที่จะส่งอีเมลหาเรา: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.ทีมงาน {{.SiteName}}<br>",
		Ko_KR: "궁금한 사항이 있으시면 메일로 알려 주세요: <a href='mailto:{{.SupportEmail}}' style='text-decoration: none; color:#2389D7;'>{{.SupportEmail}}</a>.팀 {{.SiteName}}<br>",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "store.mgo_patient.update.mobile_exists.app_error",
		Zh_CN: "手机号码已存在",
		En_US: "Mobile phone number already exists",
		Zh_HK: "手機號碼已存在",
		Vi_VN: "Số điện thoại đã tồn tại.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "store.mgo_patient.update.find_patient.app_error",
		Zh_CN: "判断手机号码是否重复失败",
		En_US: "Judging whether the mobile phone number is duplicated error",
		Zh_HK: "判斷手機號碼是否重複失敗",
		Vi_VN: "Phán xét. Số điện thoại có lặp lại thất bại",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "utils.i18n.loaded",
		Zh_CN: "从 '%v' 加载 '%v' 系统翻译信息",
		En_US: "Loading '%v' system translate information  from '%v' ",
		Zh_HK: "從 '%v' 加載 '%v' 系統翻譯資訊",
		Vi_VN: "Hệ thống thông tin từ '%v' nạp '%v' dịch",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "web.init.debug",
		Zh_CN: "初始化 web 路由",
		En_US: "Init web route",
		Zh_HK: "初始化 web 路由",
		Vi_VN: "Khởi động web Route",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    "api.templates.welcome_body.button",
		Zh_CN: "验证邮箱",
		En_US: "Verify email",
		Zh_HK: "驗證郵箱",
		Ja_JP: "認証メールボックス",
		Vi_VN: "Xác minh email",
		Ru_RU: "Подтвердить адрес эл. почты",
		Pt_BR: "Verificar o e-mail",
		Th_TH: "ยืนยันอีเมล",
		Ko_KR: "이메일 확인",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SYSTEM_ERROR,
		Zh_CN: "系统错误",
		En_US: "System error",
		Zh_HK: "系統錯誤",
		Ja_JP: "システムエラー",
		Vi_VN: "Lỗi hệ thống",
		Ru_RU: "Системная ошибка",
		Pt_BR: "Erro no sistema",
		Th_TH: "ระบบผิดพลาด",
		Ko_KR: "시스템 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SESSION_EXPIRED_ERROR,
		Zh_CN: "登录过期",
		En_US: "Login expired",
		Zh_HK: "登入過期",
		Ja_JP: "ログイン期限切れ",
		Vi_VN: "Hết thời gian đăng nhập",
		Ru_RU: "Срок действия учетной записи истек",
		Pt_BR: "Login expirado",
		Th_TH: "ล็อกอินหมดอายุ",
		Ko_KR: "로그인이 만료됨",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    JWT_PARSE_ERROR,
		Zh_CN: "token无效",
		En_US: "Token invalid",
		Zh_HK: "token無效",
		Vi_VN: "Token không hợp lệ",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    APIKEY_PARSE_ERROR,
		Zh_CN: "签名错误",
		En_US: "Signature error",
		Zh_HK: "簽名錯誤",
		Ja_JP: "署名エラー",
		Vi_VN: "Lỗi dấu hiệu",
		Ru_RU: "Ошибка подписи",
		Pt_BR: "Erro de assinatura",
		Th_TH: "ข้อผิดพลาดลายเซ็น",
		Ko_KR: "서명 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    APIKEY_IPLIMIT_ERROR,
		Zh_CN: "IP限制",
		En_US: "IP restriction",
		Zh_HK: "IP限制",
		Ja_JP: "IP制限",
		Vi_VN: "Giới hạn IP",
		Ru_RU: "ограничение IP",
		Pt_BR: "Limitação IP",
		Th_TH: "IP จำกัด",
		Ko_KR: "IP 제한",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    MISSING_USER_ERROR,
		Zh_CN: "用户不存在",
		En_US: "User does not exist",
		Zh_HK: "用戶不存在",
		Ja_JP: "ユーザーが存在しません",
		Vi_VN: "Người dùng không tồn tại",
		Ru_RU: "Пользователь не существует",
		Pt_BR: "O usuário não existe",
		Th_TH: "ไม่มีชื่อผู้ใช้งานนี้",
		Ko_KR: "존재하지 않는 사용자입니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    MISSING_INVITE_USER_ERROR,
		Zh_CN: "邀请用户不存在，请核对邀请码",
		En_US: "Invitation user does not exist, please check the invitation code",
		Zh_HK: "邀請用戶不存在，請核對邀請碼",
		Ja_JP: "招待ユーザーが存在しません。招待コードを照合してください",
		Vi_VN: "Người dùng không hiển thị khi mời, vui lòng kiểm tra mã lời mời",
		Ru_RU: "Пользователь не существовал на момент приглашения, проверьте код приглашения",
		Pt_BR: "O usuário não existe ao convidar, verifique o código de convite",
		Th_TH: "ไม่มีชื่อผู้ใช้งานนี้ขณะเชิญ กรุณาตรวจสอบรหัสบัตรเชิญ",
		Ko_KR: "존재하지 않는 사용자입니다. 초대 시 초대 코드를 확인하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    MISSING_DATA_ERROR,
		Zh_CN: "数据不存在",
		En_US: "Data does not exist",
		Zh_HK: "數據不存在",
		Ja_JP: "データが存在しません",
		Vi_VN: "Dữ liệu không tồn tại",
		Ru_RU: "Данные не существуют",
		Pt_BR: "Os dados não existem",
		Th_TH: "ไม่มีข้อมูลปรากฏอยู่",
		Ko_KR: "존재하지 않는 데이터입니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    MISSING_BANK_ERROR,
		Zh_CN: "相应的收款方式不存在，请先设置",
		En_US: "The corresponding payment method does not exist, please set it first.",
		Zh_HK: "相應的收款管道不存在，請先設定",
		Ja_JP: "対応するする受領方法が存在しません。先に設定してください",
		Vi_VN: "Phương thức thanh toán tương ứng không tồn tại, vui lòng cài đặt trước.",
		Ru_RU: "Соответствующий способ платежа не существует, сначала установите его.",
		Pt_BR: "O método de pagamento correspondente não existe, definí-lo primeiro.",
		Th_TH: "ไม่ปรากฏวิธีการชำระเงินที่สอดคล้อง กรุณากำหนดวิธีก่อน",
		Ko_KR: "해당 결제 방법은 존재하지 않습니다. 설정을 먼저 하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SQL_ERROR,
		Zh_CN: "数据库执行错误",
		En_US: "Database execution error",
		Zh_HK: "資料庫執行錯誤",
		Vi_VN: "Cơ sở dữ liệu thực hiện sai lầm.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ACCOUNT_ADDRESS_EXIST_ERROR,
		Zh_CN: "该钱包地址已存在",
		En_US: "The wallet address already exists",
		Zh_HK: "該錢包地址已存在",
		Vi_VN: "Địa chỉ tồn tại nên ví",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ACCOUNT_COIN_SYMBOL_EXIST_ERROR,
		Zh_CN: "该币种钱包已存在",
		En_US: "The currency wallet already exists",
		Zh_HK: "該幣種錢包已存在",
		Vi_VN: "Nên ví đã tồn tại.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    DEAL_OVERDUE_ERROR,
		Zh_CN: "订单已过期",
		En_US: "The order has expired",
		Zh_HK: "訂單已過期",
		Ja_JP: "注文が期限切れです",
		Vi_VN: "Hết hạn đặt hàng",
		Ru_RU: "Срок действия заказа истек",
		Pt_BR: "O pedido expirou",
		Th_TH: "คำสั่งซื้อหมดอายุ",
		Ko_KR: "주문이 만료되었습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    DEAL_OPERA_STATUS_ERROR,
		Zh_CN: "该状态订单无法执行此操作",
		En_US: "The status order cannot perform this operation",
		Zh_HK: "該狀態訂單無法執行此操作",
		Ja_JP: "この状態の注文では、この操作は実行できません",
		Vi_VN: "Do tình trạng đơn hàng nên không thể thực hiện được thao tác này",
		Ru_RU: "Эта операция не может быть выполнена для заказа с таким статусом",
		Pt_BR: "Este pedido de status não pode ser executado nesta operação",
		Th_TH: "ไม่สามารถดำเนินการคำสั่งซื้อสถานะนี้ได้",
		Ko_KR: "현재 주문 상태에서는 이 작업을 수행할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    APPEAL_OPERA_STATUS_ERROR,
		Zh_CN: "该状态申诉无法执行此操作",
		En_US: "This status appeal cannot perform this operation",
		Zh_HK: "該狀態申訴無法執行此操作",
		Ja_JP: "この状態の申立てでは、この操作は実行できません",
		Vi_VN: "Do tình trạng khiếu nại nên không thể thực hiện được thao tác này",
		Ru_RU: "Эта операция не может быть выполнена для жалобы с таким статусом",
		Pt_BR: "Esta reclamação de status não pode ser executada nesta operação",
		Th_TH: "ไม่สามารถดำเนินการคำร้องเรียนสถานะนี้ได้",
		Ko_KR: "현재 불만사항 상태에서는 이 작업을 수행할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    OPERA_STATUS_ERROR,
		Zh_CN: "无法执行此操作",
		En_US: "Cannot perform this operation",
		Zh_HK: "無法執行此操作",
		Ja_JP: "この操作は実行できません",
		Vi_VN: "Không thể thực hiện thao tác này",
		Ru_RU: "Невозможно выполнить эту операцию",
		Pt_BR: "Não é possível executar esta operação",
		Th_TH: "ไม่สามารถดำเนินการเรื่องนี้ได้",
		Ko_KR: "이 작업을 수행할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    OPERA_PERMISSION_DENIED_ERROR,
		Zh_CN: "没有权限",
		En_US: "Permission denied",
		Zh_HK: "沒有許可權",
		Ja_JP: "許可権がありません",
		Vi_VN: "Quyền đã bị từ chối",
		Ru_RU: "В разрешении отказано",
		Pt_BR: "Permissão negada",
		Th_TH: "การอนุญาตถูกปฏิเสธ",
		Ko_KR: "권한 거부됨",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    AVAILABLE_BALANCE_NOT_ENOUGH_ERROR,
		Zh_CN: "可用余额不足",
		En_US: "Insufficient available balance",
		Zh_HK: "可用餘額不足",
		Ja_JP: "残高不足です",
		Vi_VN: "Không đủ số dư trong tài khoản",
		Ru_RU: "Недостаточный баланс счета",
		Pt_BR: "Saldo disponível insuficiente",
		Th_TH: "ยอดเงินคงเหลือไม่พอ",
		Ko_KR: "부족한 사용 가능 잔액",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ITEM_ESTABLISH_ERROR,
		Zh_CN: "该广告无法下单",
		En_US: "The advertisement could not be placed",
		Zh_HK: "該廣告無法下單",
		Ja_JP: "この広告は発注できません",
		Vi_VN: "Đơn hàng này không thể đặt cho quảng cáo.",
		Ru_RU: "Невозможно разместить заказ для этой рекламы.",
		Pt_BR: "O pedido não pode ser colocado para o anúncio.",
		Th_TH: "ไม่สามารถวางคำสั่งซื้อสำหรับโฆษณาได้",
		Ko_KR: "해당 광고에 대한 주문이 발주될 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ITEM_ESTABLISH_SELF_ERROR,
		Zh_CN: "自己的广告无法下单",
		En_US: "You can't place an order for your advertisement",
		Zh_HK: "自己的廣告無法下單",
		Ja_JP: "自分の広告が発注できません",
		Vi_VN: "Đơn hàng này không thể đặt cho quảng cáo của bạn.",
		Ru_RU: "Невозможно разместить заказ для вашей рекламы.",
		Pt_BR: "O pedido não pode ser colocado para o seu anúncio.",
		Th_TH: "ไม่สามารถวางคำสั่งซื้อสำหรับโฆษณาของคุณได้",
		Ko_KR: "귀하의 광고에 대한 주문이 발주될 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    AMOUNT_OUT_LIMIT_ERROR,
		Zh_CN: "数量超出限制范围",
		En_US: "Quantity beyond the limit",
		Zh_HK: "數量超出限制範圍",
		Ja_JP: "数量が制限範囲を超えています",
		Vi_VN: "Số lượng vượt quá giới hạn",
		Ru_RU: "Количество превышает лимит",
		Pt_BR: "A quantidade excede o limite",
		Th_TH: "ปริมาณเกินขีดจำกัด",
		Ko_KR: "수량이 한도를 초과하였습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    VERIFIED_CODE_ERROR,
		Zh_CN: "验证码错误",
		En_US: "Verification code error",
		Zh_HK: "驗證碼錯誤",
		Ja_JP: "認証コードエラー",
		Vi_VN: "Mã xác minh có lỗi",
		Ru_RU: "Ошибка кода подтверждения",
		Pt_BR: "Erro no código de verificação",
		Th_TH: "รหัสยืนยันมีข้อผิดพลาด",
		Ko_KR: "인증 코드 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_PWD_VALID_ERROR,
		Zh_CN: "密码格式错误(6位以上含大小写字母和数字)",
		En_US: "Password Format Error(6 or more with uppercase and lowercase letters and numbers)",
		Zh_HK: "密碼格式錯誤(6位以上含大小寫字母和數字)",
		Ja_JP: "パスワード形式エラー",
		Vi_VN: "Định dạng mật khẩu không đúng",
		Ru_RU: "Неверный формат пароля",
		Pt_BR: "Formato de senha errada ",
		Th_TH: "รูปแบบรหัสผ่านไม่ถูกต้อง",
		Ko_KR: "잘못된 비밀번호 형식",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_FUND_PWD_VALID_ERROR,
		Zh_CN: "密码格式错误",
		En_US: "Password Format Error",
		Zh_HK: "密碼格式錯誤",
		Ja_JP: "パスワード形式エラー",
		Vi_VN: "Định dạng mật khẩu không đúng",
		Ru_RU: "Неверный формат пароля",
		Pt_BR: "Formato de senha errada ",
		Th_TH: "รูปแบบรหัสผ่านไม่ถูกต้อง",
		Ko_KR: "잘못된 비밀번호 형식",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_LOGIN_BLANK_PWD_ERROR,
		Zh_CN: "密码错误",
		En_US: "Password error",
		Zh_HK: "密碼錯誤",
		Ja_JP: "パスワードエラー",
		Vi_VN: "Mật khẩu không đúng",
		Ru_RU: "Неверный пароль",
		Pt_BR: "Senha errada",
		Th_TH: "รหัสผ่านไม่ถูกต้อง",
		Ko_KR: "비밀번호 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_LOGIN_INACTIVE_ERROR,
		Zh_CN: "账号被禁用",
		En_US: "Account Banned",
		Zh_HK: "帳號被禁用",
		Ja_JP: "アカウントが無効です",
		Vi_VN: "Tài khoản bị vô hiệu hóa",
		Ru_RU: "Лицевой счет отключен",
		Pt_BR: "A conta está desativada",
		Th_TH: "บัญชีถูกยกเลิก",
		Ko_KR: "계정이 비활성화되었습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_LOGIN_ATTEMPTS_TOO_MANY_ERROR,
		Zh_CN: "错误次数太多,请 {{.Name}}s 后再试",
		En_US: "There are too many error. Please try again {{.Name}}s later",
		Zh_HK: "錯誤次數太多，請 {{.Name}}s 再試",
		Ja_JP: "エラー回数が多すぎます。 請 {{.Name}} 秒後にリトライしてください",
		Vi_VN: "Xảy ra quá nhiều lỗi, vui lòng thử lại sau {{.Name}} giây",
		Ru_RU: "Слишком много ошибок, попробуйте снова через {{.Name}} с",
		Pt_BR: "Muitos erros, tentar novamente depois de {{.Name}} segundos",
		Th_TH: "เกิดข้อผิดพลาดมากเกินไป กรุณาลองใหม่ในอีก {{.Name}} วินาที",
		Ko_KR: "허용된 실패 횟수가 초과되었습니다. {{.Name}} 초 후에 다시 시도하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    VERIFICATION_CODE_SEND_FREQUENTLY_ERROR,
		Zh_CN: "验证码发送太频繁，{{ .Name }}s后再试",
		En_US: "Verification code sent too frequently，retry after {{ .Name }}s",
		Zh_HK: "驗證碼發送太頻繁，{{ .Name }}s後再試",
		Ja_JP: "認証コードの送信が多すぎます。{{ .Name }}秒後にリトライしてください",
		Vi_VN: "Mã xác minh bị gửi quá nhiều lần liên tiếp, hãy thử lại sau {{ .Name }} giây",
		Ru_RU: "Проверочный код отправлялся слишком часто, попробуйте снова через {{ .Name }} с",
		Pt_BR: "O código de verificação foi enviado com muita frequência, tente novamente depois de {{ .Name }} segundos",
		Th_TH: "ส่งรหัสยืนยันถี่เกินไป ลองใหม่ในอีก {{ .Name }} วินาที",
		Ko_KR: "인증 코드가 너무 자주 전달되었습니다. {{ .Name }} 초 후에 다시 시도하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SEND_VERIFY_ERROR,
		Zh_CN: "验证码发送失败",
		En_US: "Verification Code failed to send",
		Zh_HK: "驗證碼發送失敗",
		Ja_JP: "認証コードの送信に失敗しました",
		Vi_VN: "Không thể gửi mã xác minh",
		Ru_RU: "Не удалось отправить проверочный код",
		Pt_BR: "Houve uma falha ao enviar o código de verificação",
		Th_TH: "ส่งรหัสยืนยันล้มเหลว",
		Ko_KR: "인증 코드 전달에 실패하였습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SEND_VERIFY_NONE_AVAILABLE_ERROR,
		Zh_CN: "地址不可用，验证码发送失败",
		En_US: "Address unavailable, authentication code failed to send",
		Zh_HK: "地址不可用，驗證碼發送失敗",
		Ja_JP: "認証コードの送信に失敗しました",
		Vi_VN: "Không thể gửi mã xác minh",
		Ru_RU: "Не удалось отправить проверочный код",
		Pt_BR: "Houve uma falha ao enviar o código de verificação",
		Th_TH: "ส่งรหัสยืนยันล้มเหลว",
		Ko_KR: "인증 코드 전달에 실패하였습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_TOO_LARGE_ERROR,
		Zh_CN: "文件太大",
		En_US: "File is too large",
		Zh_HK: "檔案太大",
		Ja_JP: "ファイルが大きすぎます",
		Vi_VN: "Dung lượng tập tin quá lớn",
		Ru_RU: "Размер файла слишком велик",
		Pt_BR: "Tamanho do arquivo é muito grande",
		Th_TH: "ขนาดไฟล์ใหญ่เกินไป",
		Ko_KR: "파일 크기가 너무 큽니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_PARSE_ERROR,
		Zh_CN: "文件解析错误",
		En_US: "File parsing error",
		Zh_HK: "檔案解析錯誤",
		Ja_JP: "ファイル解析エラー",
		Vi_VN: "Lỗi phân tích tập tin",
		Ru_RU: "Ошибка синтаксического разбора файла",
		Pt_BR: "Erro de análise do arquivo",
		Th_TH: "กระบวนการตรวจสอบโครงสร้างไฟล์มีความผิดพลาด",
		Ko_KR: "파일 구문 분석 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_NAME_ERROR,
		Zh_CN: "文件名不合法",
		En_US: "Illegal filename",
		Zh_HK: "檔名不合法",
		Ja_JP: "ファイル名は不正です",
		Vi_VN: "Tên tập tin lậu",
		Ru_RU: "Неверное имя файла",
		Pt_BR: "Nome do ficheiro ilegal",
		Th_TH: "ชื่อไฟล์ไม่ถูกต้อง",
		Ko_KR: "파일 이름 이 비합법적 이다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_NO_FIND_ERROR,
		Zh_CN: "请上传文件",
		En_US: "Please upload the file",
		Zh_HK: "請上傳文件",
		Ja_JP: "ファイルをアップロードしてください",
		Vi_VN: "Vui lòng tải tập tin lên",
		Ru_RU: "Выгрузите файл",
		Pt_BR: "Carregar um arquivo",
		Th_TH: "กรุณาอัปโหลดไฟล์",
		Ko_KR: "파일을 업로드하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_SAVE_ERROR,
		Zh_CN: "文件保存错误",
		En_US: "File saving error",
		Zh_HK: "檔案保存錯誤",
		Ja_JP: "ファイル保存エラー",
		Vi_VN: "Lỗi lưu tập tin",
		Ru_RU: "Ошибка сохранения файла",
		Pt_BR: "Erro ao salvar arquivo",
		Th_TH: "การบันทึกไฟล์มีความผิดพลาด",
		Ko_KR: "파일 저장 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FILE_READ_ERROR,
		Zh_CN: "文件读取错误",
		En_US: "File read error",
		Zh_HK: "檔案讀取錯誤",
		Vi_VN: "Lỗi đọc tập tin",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    OPERA_REPEAT_ERROR,
		Zh_CN: "重复操作",
		En_US: "Repetitive operation",
		Zh_HK: "重複操作",
		Ja_JP: "繰り返し操作です",
		Vi_VN: "Lặp lại thao tác",
		Ru_RU: "Повторите действие",
		Pt_BR: "Repetir operação",
		Th_TH: "ดำเนินการซ้ำ",
		Ko_KR: "반복 작업",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_EXIST_ERROR,
		Zh_CN: "该邮箱已注册",
		En_US: "The email registered",
		Zh_HK: "該郵箱已註冊",
		Ja_JP: "そのメールボックスは登録済みです",
		Vi_VN: "Hộp thư đã được đăng ký",
		Ru_RU: "Почтовый ящик зарегистрирован",
		Pt_BR: "A caixa de correio foi registrada",
		Th_TH: "ลงทะเบียนกล่องข้อความแล้ว",
		Ko_KR: "메일함이 등록되었습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USERNAME_EXIST_ERROR,
		Zh_CN: "用户名已存在",
		En_US: "User name already exists",
		Zh_HK: "用戶名已存在",
		Ja_JP: "ユーザー名は存在しています",
		Vi_VN: "Tên người dùng này đã tồn tại",
		Ru_RU: "Имя пользователя уже существует",
		Pt_BR: "O nome de usuário já existe",
		Th_TH: "มีชื่อผู้ใช้งานนี้แล้ว",
		Ko_KR: "이미 사용 중인 사용자 이름입니다.",
	})
	i18ns.AddI18n(utils.I18n{
		Id:    NICKNAME_EXIST_ERROR,
		Zh_CN: "昵称已存在",
		En_US: "Nick name already exists",
		Zh_HK: "昵稱已存在",
		Ja_JP: "ニックネームは存在しています",
		Vi_VN: "Biệt danh này đã tồn tại",
		Ru_RU: "Псевдоним уже существует",
		Pt_BR: "O apelido já existe",
		Th_TH: "มีชื่อเล่นนี้แล้ว",
		Ko_KR: "이미 사용 중인 닉네임입니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    WHITE_BLACK_ADD_SELF_ERROR,
		Zh_CN: "不能添加自己",
		En_US: "Can't add yourself",
		Zh_HK: "不能添加自己",
		Vi_VN: "Không thể thêm tự",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    PASS_DEAL_MESSAGE,
		Zh_CN: "已付款，请放行",
		En_US: "Payment made, please let it go",
		Zh_HK: "已付款，請放行",
		Ja_JP: "支払い済みです。リリースしてください",
		Vi_VN: "Đã thanh toán, vui lòng phát hành",
		Ru_RU: "Оплачено, пожалуйста, выпустите",
		Pt_BR: "Pago, liberar",
		Th_TH: "ชำระเงินแล้ว กรุณาปล่อยเหรียญ",
		Ko_KR: "결제되었습니다. 출금 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    NO_COMPLETE_DEAL,
		Zh_CN: "有未完成的订单，无法下架",
		En_US: "Unfinished orders, Unable to get off the shelf",
		Zh_HK: "有未完成的訂單，無法下架",
		Ja_JP: "未完成の注文があります。取り下げることができません",
		Vi_VN: "Một số đơn hàng chưa thực hiện xong nên không thể gỡ bỏ",
		Ru_RU: "Имеются незавершенные заказы, которые нельзя удалить",
		Pt_BR: "Existem pedidos inacabados que não podem ser removidos",
		Th_TH: "มีคำสั่งซื้อที่ไม่เสร็จสมบูรณ์ซึ่งไม่สามารถลบทิ้งได้",
		Ko_KR: "마감되지 않은 주문내역이 있어 제거할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    APPEAL_DEAL_MESSAGE,
		Zh_CN: "有申诉，请及时处理",
		En_US: "If there is a complaint, please deal with it in time",
		Zh_HK: "有申訴，請及時處理",
		Ja_JP: "申立てがあります。速やかに処理してください",
		Vi_VN: "Nếu có khiếu nại, vui lòng xử lý kịp thời",
		Ru_RU: "Имеется жалоба, пожалуйста, обработайте ее своевременно",
		Pt_BR: "Tem uma reclamação, lidar com ela a tempo",
		Th_TH: "มีคำร้องเรียน กรุณาจัดการภายในเวลาที่กำหนด",
		Ko_KR: "불만사항이 제기되었습니다. 지정된 시간 내에 처리하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    APPEAL_DEAL_MESSAGE2,
		Zh_CN: "该订单被申诉 {{ .Name }}，交易暫時被冻结，如有疑問，請與在線客服聯繫。",
		En_US: "The order was appealed {{ .Name }} and the transaction was temporarily frozen. If in doubt, please contact the online customer service.",
		Zh_HK: "該訂單被申訴 {{ .Name }}，交易暫時被凍結，如有疑問，請與線上客服聯系。",
		Ja_JP: "この注文書は {{ .Name }} として提出されました。取引は一時的に凍結されました。疑問があれば、オンラインカスタマーサービスに連絡してください。",
		Vi_VN: "Đơn đặt hàng đã được kêu {{ .Name }} và giao dịch bị tạm thời đóng băng. Nếu bị nghi ngờ, hãy liên lạc với dịch vụ khách hàng trực tuyến.",
		Ru_RU: "этот заказ был обжалован {{ .Name }} сделки были временно заморожены, и в случае сомнений просьба обращаться в интернет - сервис.",
		Pt_BR: "A encomenda FOI invocada {{ .Name }} e a transacção FOI temporariamente congelada. Em CaSO de dúvida, por favor contacte o serviço de clientes online.",
		Th_TH: "คำสั่งนี้ได้รับการร้องเรียนและรายการที่ถูกระงับชั่วคราวหากมีข้อสงสัยโปรดติดต่อฝ่ายบริการออนไลน์",
		Ko_KR: "이 주문서는 고소 {{ .Name }}, 거래가 잠시 동결되었는데, 의문이 있으면 온라인 고객복과 연락해 주세요.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BUY_DEAL_MESSAGE,
		Zh_CN: "有用户下单，请及时处理",
		En_US: "If there is an order from a user, please deal with it in time",
		Zh_HK: "有用戶下單，請及時處理",
		Ja_JP: "ユーザーが発注しています。速やかに処理してください",
		Vi_VN: "Một số người dùng đã đặt hàng, vui lòng xử lý kịp thời.",
		Ru_RU: "Некоторые пользователи разместили заказы, пожалуйста, обработайте их своевременно.",
		Pt_BR: "Alguns usuários fizeram pedidos, lidar com eles a tempo.",
		Th_TH: "ผู้ใช้งานบางรายวางคำสั่งซื้อ กรุณาจัดการภายในเวลาที่กำหนด",
		Ko_KR: "일부 사용자가 주문을 신청하였습니다. 지정된 시간 내에 처리하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    WITHDRAW_SUCCESS_MESSAGE,
		Zh_CN: "提币成功",
		En_US: "Successful withdrawal of currency",
		Zh_HK: "提幣成功",
		Ja_JP: "出金成功",
		Vi_VN: "Rút tiền thành công",
		Ru_RU: "Вывод завершен",
		Pt_BR: "Retirada bem sucedida",
		Th_TH: "ถอนเงินสำเร็จ",
		Ko_KR: "성공적으로 출금하였습니다.",
	})
	i18ns.AddI18n(utils.I18n{
		Id:    WALLET_HTTP_ERROR,
		Zh_CN: "钱包服务错误",
		En_US: "Wallet service error",
		Zh_HK: "錢包服務錯誤",
		Ja_JP: "ウォレットサービスエラー",
		Vi_VN: "Lỗi dịch vụ ví",
		Ru_RU: "Ошибка обслуживания кошелька",
		Pt_BR: "Erro no serviço da carteira",
		Th_TH: "บริการกระเป๋าเงินเกิดความผิดพลาด",
		Ko_KR: "지갑 서비스 오류",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    REQUEST_RATE_LIMIT,
		Zh_CN: "请求太频繁",
		En_US: "Requests are too frequent",
		Zh_HK: "請求太頻繁",
		Vi_VN: "Yêu cầu quá thường xuyên.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    FEE_NOT_ENOUGH,
		Zh_CN: "手续费不足",
		En_US: "Insufficient handling fees",
		Zh_HK: "手續費不足",
		Ja_JP: "手数料不足",
		Vi_VN: "Phí không đủ",
		Ru_RU: "Недостаточные комиссионные",
		Pt_BR: "Taxas insuficientes",
		Th_TH: "ค่าธรรมเนียมไม่พอ",
		Ko_KR: "수수료 부족",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    AMOUNT_NOT_ENOUGH,
		Zh_CN: "提币数不足最小提币",
		En_US: "Less than the minimum withdrawal",
		Zh_HK: "提幣數不足最小提幣",
		Vi_VN: "Chưa đến mức thu hồi tối thiểu",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_USERNAME_ERROR,
		Zh_CN: "用户名不合法",
		En_US: "Illegal username",
		Zh_HK: "用戶名不合法",
		Ja_JP: "ユーザー名が規則に合致していません",
		Vi_VN: "Tên người dùng không hợp pháp",
		Ru_RU: "Недопустимое имя пользователя",
		Pt_BR: "O nome de usuário é ilegal",
		Th_TH: "ชื่อผู้ใช้งานผิดกฎหมาย",
		Ko_KR: "사용자 이름을 사용할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    USER_NICKNAME_ERROR,
		Zh_CN: "昵称不合法",
		En_US: "Illegal nickname",
		Zh_HK: "昵稱不合法",
		Ja_JP: "ニックネームが規則に合致していません",
		Vi_VN: "Biệt danh không hợp pháp",
		Ru_RU: "Недопустимый псевдоним",
		Pt_BR: "O apelido é ilegal",
		Th_TH: "ชื่อเล่นผิดกฎหมาย",
		Ko_KR: "닉네임을 사용할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOT_VERIFIED,
		Zh_CN: "E-mail尚未验证",
		En_US: "E-mail not yet validated",
		Zh_HK: "E-mail尚未驗證",
		Ja_JP: "E-mailが未認証です",
		Vi_VN: "Email chưa được xác minh",
		Ru_RU: "Адрес эл. почты еще не подтвержден",
		Pt_BR: "O e-mail ainda não foi verificado",
		Th_TH: "ยังไม่ได้ยืนยันอีเมล",
		Ko_KR: "이메일이 아직 확인되지 않았습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_ORDER,
		Zh_CN: "订单编号",
		En_US: "Order number",
		Zh_HK: "訂單編號",
		Ja_JP: "注文番号",
		Vi_VN: "Mã số đơn hàng",
		Ru_RU: "Номер заказа",
		Pt_BR: "Número do pedido",
		Th_TH: "หมายเลขคำสั่งซื้อขาย",
		Ko_KR: "주문 번호",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_ITEM,
		Zh_CN: "广告编号",
		En_US: "Advertising Number",
		Zh_HK: "廣告編號",
		Ja_JP: "広告番号",
		Vi_VN: "Mã số quảng cáo",
		Ru_RU: "Номер рекламы",
		Pt_BR: "Número do anúncio",
		Th_TH: "หมายเลขโฆษณา",
		Ko_KR: "광고 번호",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_HELLO,
		Zh_CN: "尊敬的 {{ .Name }} 您好：",
		En_US: "Dear  {{ .Name }} Hello:",
		Zh_HK: "尊敬的 {{ .Name }} 您好:",
		Ja_JP: "敬愛する {{ .Name }} さま、こんにちは、",
		Vi_VN: "Kính gửi {{ .Name }}, Xin chào:",
		Ru_RU: "Здравствуйте, уважаемый {{ .Name }}!",
		Pt_BR: "Caro {{ .Name }}, olá:",
		Th_TH: "สวัสดีคุณ {{ .Name }} ",
		Ko_KR: "{{ .Name }}님께, 안녕하세요.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_INTRODUCE1,
		Zh_CN: "GIAC为您提供安全，可信赖的数字资产交易及资产管理服务！",
		En_US: "GIAC provides you with safe and reliable digital asset trading and asset management services!",
		Zh_HK: "GIAC為您提供安全，可信賴的數位資產交易及資產管理服務！",
		Ja_JP: "GIACは、安全で信頼できるデジタル資産取引と資産管理サービスを提供します。",
		Vi_VN: "GIAC cung cấp các dịch vụ quản lý tài sản và giao dịch tài sản kỹ thuật số an toàn, đáng tin cậy!",
		Ru_RU: "GIAC предоставляет Вам безопасные и надежные службы для торговли и управления цифровыми активами!",
		Pt_BR: "A GIAC oferece a você serviços de gerenciamento de ativo e negociação de ativos digitais seguros e confiáveis!",
		Th_TH: "GIAC ให้บริการซื้อขายสกุลเงินดิจิทัลและการบริหารทรัพย์สินที่ปลอดภัยและไว้วางใจได้",
		Ko_KR: "GIAC은 안전하고 신뢰할 수 있는 전자 자산 거래와 자산 관리 서비스를 제공합니다!",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_INTRODUCE2,
		Zh_CN: "感谢您在GIAC完成了USDT币的交易，您可以随时登入您在GIAC APP注册的账户，查询您的订单详情。",
		En_US: "Thank you for completing the purchase of USDT coins in GIAC. You can log in to your account registered with GIAC APP at any time and inquire about the details of your order.",
		Zh_HK: "感謝您在GIAC完成了USDT幣的交易，您可以隨時登入您在GIAC APP註冊的帳戶，査詢您的訂單詳情。",
		Ja_JP: "GIACで、GIACの取引を完了させてくださったことに感謝いたします。",
		Vi_VN: "Cảm ơn bạn đã hoàn thành giao dịch tiền xu GIAC tại GIAC, bạn có thể đăng nhập vào tài khoản mà bạn đã đăng ký trên ứng dụng của evertoken bất cứ lúc nào để kiểm tra chi tiết đơn đặt hàng của bạn.",
		Ru_RU: "Благодарим за выполнение денежной транзакции в GIAC. Вы можете в любое время войти в свой лицевой счет, зарегистрированный в приложении GIAC, чтобы ознакомиться с подробностями по своему заказу.",
		Pt_BR: "Obrigado por concluir a transação com moedas GIAC na GIAC. Você pode fazer login na conta que registrou no app Everytoken a qualquer momento para verificar os detalhes de seu pedido.",
		Th_TH: "ขอบคุณที่ทำธุรกรรมซื้อขายเหรียญ GIAC ที่ GIAC คุณสามารถล็อกอินเข้าสู่บัญชีที่คุณลงทะเบียนไว้ในแอปฯ GIAC เพื่อตรวจดูรายละเอียดคำสั่งซื้อขายของคุณได้ทุกเวลา",
		Ko_KR: "GIAC에서 GIAC 코인 거래를 완료해 주셔서 감사합니다. GIAC 어플에 등록해 주신 계정으로 로그인하여 언제든지 귀하의 주문 세부내역을 확인하실 수 있습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_ORDERT1,
		Zh_CN: "购买 {{.SYMBOL}}",
		En_US: "BUY {{.SYMBOL}}",
		Zh_HK: "購買 {{.SYMBOL}}",
		Ja_JP: "購入 {{.SYMBOL}}",
		Vi_VN: "Mua {{.SYMBOL}}",
		Ru_RU: "Купить {{.SYMBOL}}",
		Pt_BR: "Comprar {{.SYMBOL}}",
		Th_TH: "ซื้อ {{.SYMBOL}}",
		Ko_KR: "구매 {{.SYMBOL}}",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_ORDERT2,
		Zh_CN: "出售 {{.SYMBOL}}",
		En_US: "SELL {{.SYMBOL}}",
		Zh_HK: "出售 {{.SYMBOL}}",
		Ja_JP: "売却 {{.SYMBOL}}",
		Vi_VN: "Bán {{.SYMBOL}}",
		Ru_RU: "Продать {{.SYMBOL}}",
		Pt_BR: "Vender {{.SYMBOL}}",
		Th_TH: "ขาย {{.SYMBOL}}",
		Ko_KR: "판매 {{.SYMBOL}}",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_BUSINESST,
		Zh_CN: "商家",
		En_US: "Business",
		Zh_HK: "商家",
		Ja_JP: "ショップ",
		Vi_VN: "Bên bán",
		Ru_RU: "Продавец",
		Pt_BR: "Comerciante",
		Th_TH: "ผู้ค้า",
		Ko_KR: "판매자",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_AMOUNTT,
		Zh_CN: "订单金额",
		En_US: "Amount",
		Zh_HK: "訂單金額",
		Ja_JP: "注文金額",
		Vi_VN: "Giá trị đơn hàng",
		Ru_RU: "Сумма заказа",
		Pt_BR: "Quantidade do pedido",
		Th_TH: "จำนวนที่สั่ง",
		Ko_KR: "주문량",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_PRICET,
		Zh_CN: "单价",
		En_US: "Price",
		Zh_HK: "單價",
		Ja_JP: "単価",
		Vi_VN: "Đơn giá",
		Ru_RU: "Цена за единицу товара",
		Pt_BR: "Preço unitário",
		Th_TH: "ราคาต่อหน่วย",
		Ko_KR: "단가",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_QUANTITYT,
		Zh_CN: "数量",
		En_US: "Quantity",
		Zh_HK: "數量",
		Ja_JP: "数量",
		Vi_VN: "Số lượng",
		Ru_RU: "Количество",
		Pt_BR: "Quantidade",
		Th_TH: "ปริมาณ",
		Ko_KR: "수량",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_METHODT,
		Zh_CN: "支付方式",
		En_US: "Payment",
		Zh_HK: "支付方式",
		Ja_JP: "支払い方法",
		Vi_VN: "Phương pháp thanh toán",
		Ru_RU: "Способ оплаты",
		Pt_BR: "Método de pagamento",
		Th_TH: "วิธีชำระเงิน",
		Ko_KR: "결제 방법",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_PAYCODET,
		Zh_CN: "付款参考号",
		En_US: "Payment reference number",
		Zh_HK: "付款參考號",
		Ja_JP: "支払いレファレンスナンバー",
		Vi_VN: "Số tham chiếu thanh toán",
		Ru_RU: "Код платежа",
		Pt_BR: "Número de referência do pagamento",
		Th_TH: "หมายเลขอ้างอิงการชำระเงิน",
		Ko_KR: "결제 참조 번호",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_CREATEATT,
		Zh_CN: "下单时间",
		En_US: "Order time",
		Zh_HK: "下單時間",
		Ja_JP: "発注時間",
		Vi_VN: "Thời gian đặt hàng",
		Ru_RU: "Время для размещения заказа",
		Pt_BR: "Hora de fazer um pedido",
		Th_TH: "ช่วงเวลาวางคำสั่งซื้อขาย",
		Ko_KR: "주문 시간",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUST,
		Zh_CN: "订单状态",
		En_US: "Order status",
		Zh_HK: "訂單狀態",
		Ja_JP: "注文状態",
		Vi_VN: "Trạng thái đơn hàng",
		Ru_RU: "Состояние заказа",
		Pt_BR: "Status do pedido",
		Th_TH: "สถานะคำสั่งซื้อขาย",
		Ko_KR: "주문 상태",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_SUCCESS,
		Zh_CN: "成功",
		En_US: "Success",
		Zh_HK: "成功",
		Ja_JP: "成功",
		Vi_VN: "Thành công",
		Ru_RU: "Успешно",
		Pt_BR: "Sucesso",
		Th_TH: "สำเร็จ",
		Ko_KR: "성공",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_WAITING,
		Zh_CN: "已支付",
		En_US: "Paid",
		Zh_HK: "已支付",
		Ja_JP: "支払い済",
		Vi_VN: "Đã thanh toán",
		Ru_RU: "Оплачено",
		Pt_BR: "Pago",
		Th_TH: "ชำระเงินแล้ว",
		Ko_KR: "결제 완료됨",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_APPEAL,
		Zh_CN: "申诉中",
		En_US: "In complaint",
		Zh_HK: "申訴中",
		Ja_JP: "申立て中",
		Vi_VN: "Đang trong quá trình khiếu nại",
		Ru_RU: "В процессе подачи жалобы",
		Pt_BR: "No processo de reclamação",
		Th_TH: "อยู่ในขั้นตอนร้องเรียน",
		Ko_KR: "불만사항이 프로세스 중에 있습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_SOLD,
		Zh_CN: "已完成",
		En_US: "Completed",
		Zh_HK: "已完成",
		Ja_JP: "完了",
		Vi_VN: "Đã xong",
		Ru_RU: "Выполнено",
		Pt_BR: "Concluído",
		Th_TH: "เสร็จสิ้น",
		Ko_KR: "완료됨",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_NOTSALE,
		Zh_CN: "已下架",
		En_US: "Off-shelf",
		Zh_HK: "已下架",
		Ja_JP: "取り下げ済み",
		Vi_VN: "Đã được gỡ bỏ",
		Ru_RU: "Удалено",
		Pt_BR: "Foi removido",
		Th_TH: "ลบออกแล้ว",
		Ko_KR: "제거되었습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_ACTIVE,
		Zh_CN: "等待付款",
		En_US: "Waiting for payment",
		Zh_HK: "等待付款",
		Ja_JP: "支払いを待つ",
		Vi_VN: "Đợi thanh toán",
		Ru_RU: "ждать уплаты",
		Pt_BR: "Esperando o pagamento",
		Th_TH: "รอการชำระเงิน",
		Ko_KR: "지불을 기다리다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_STATUS_CANCEL,
		Zh_CN: "已取消",
		En_US: "Cancelled",
		Zh_HK: "已取消",
		Ja_JP: "キャンセル",
		Vi_VN: "Đã hủy",
		Ru_RU: "Отменено",
		Pt_BR: "Cancelado",
		Th_TH: "ยกเลิกแล้ว",
		Ko_KR: "취소됨",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTET,
		Zh_CN: "重要说明",
		En_US: "Important note",
		Zh_HK: "重要說明",
		Ja_JP: "大事な説明",
		Vi_VN: "Lưu ý quan trọng",
		Ru_RU: "Важное примечание",
		Pt_BR: "Nota importante",
		Th_TH: "หมายเหตุที่สำคัญ",
		Ko_KR: "중요한 공지사항",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTE1,
		Zh_CN: "您之所以收到这封邮件，是因为您曾经注册成为GIAC的用户。",
		En_US: "You received this email because you have registered as a GIAC user.",
		Zh_HK: "您之所以收到這封郵件，是因為您曾經注册成為GIAC的用戶。",
		Ja_JP: "お客さまが以前GIACのユーザーであったことから、このメールを受信しています。",
		Vi_VN: "Bạn nhận được email này vì bạn đã đăng ký làm người dùng của GIAC.",
		Ru_RU: "Вы получили это письмо в ответ на регистрацию в системе GIAC.",
		Pt_BR: "O motivo pelo qual você recebeu este e-mail é que você se registrou como um usuário da GIAC.",
		Th_TH: "คุณได้รับอีเมลฉบับนี้เนื่องจากคุณได้ลงทะเบียนเป็นผู้ใช้งานของ GIAC",
		Ko_KR: "GIAC에 사용자로 등록해 주셔서 본 이메일이 귀하에게 송부되었습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTE2,
		Zh_CN: "GIAC是点对点的商品流通数字货币交易平台，为百万用户提供安全,可信赖的数字资产管理服务。",
		En_US: "GIAC is a point-to-point digital currency trading platform for commodity circulation, providing secure and reliable digital asset management services for millions of users.",
		Zh_HK: "GIAC是點對點的商品流通數位貨幣交易平臺，為百萬用戶提供安全，可信賴的數位資產管理服務。",
		Ja_JP: "GIACポイント・ツー・ポイントの商品流通デジタル通貨取引プラットフォームでは、100万人のユーザーのために安全で信頼性のあるデジタル資産管理サービスを提供します。",
		Vi_VN: "GIAC là một nền tảng giao dịch tiền số lưu thông hàng hóa theo giao thức kết nối các điểm, cung cấp dịch vụ quản lý tài sản số an toàn, đáng tin cậy cho hàng triệu người dùng.",
		Ru_RU: "GIAC — торговая площадка для двустороннего товарообмена цифровыми валютами, предоставляющая безопасные и надежные службы управления цифровыми активами для миллионов пользователей.",
		Pt_BR: "A GIAC é uma plataforma de negociação de moeda digital de circulação de bens ponto a ponto que fornece serviços de gerenciamento de ativos digitais seguros e confiáveis para milhões de usuários.",
		Th_TH: "GIAC คือแพลตฟอร์มสำหรับซื้อขายสกุลเงินดิจิทัลที่มีการหมุนเวียนผลิตภัณฑ์แบบจุดต่อจุด โดยให้บริการบริหารทรัพย์สินดิจิทัลที่ปลอดภัยและไว้วางใจได้แก่ผู้ใช้งานหลายล้านราย",
		Ko_KR: "GIAC은 포인트투포인트 기술로 이루어진 상품 순환 전자 화폐 거래 플랫폼으로 안전하고 신뢰할 수 있는 전자 자산 관리 서비스를 수백만 명의 사용자에게 제공하고 있습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTE3,
		Zh_CN: "重要提醒：请在GIAC币的交易过程中，核实交易及支付信息，谨防信息错误导致资产损失。",
		En_US: "Important Reminder: Please verify the transaction and payment information in the GIAC currency transaction process, and guard against the loss of assets caused by information errors.",
		Zh_HK: "重要提醒：請在GIAC幣的交易過程中，核實交易及支付資訊，謹防資訊錯誤導致資產損失。",
		Ja_JP: "重要事項：GIAC通貨の取引過程においては、取引と支払い情報を照合確認することで、情報エラーによる資産の損失を厳格に防ぎます。",
		Vi_VN: "Lưu ý quan trọng: Vui lòng kiểm tra thông tin giao dịch và thanh toán trong quá trình giao dịch tiền tệ với GIAC, hãy cẩn luôn cẩn trọng với các loại thông tin.",
		Ru_RU: "Важное напоминание: Проверяйте всю информацию по операции и оплате при проведении транзакции с валютой GIAC, будьте осторожны.Ошибки ведут к потере активов.",
		Pt_BR: "Lembrete importante: Verificar as informações de transação e de pagamento durante a transação da moeda GIAC, cuidado com as informações.",
		Th_TH: "คำเตือนสำคัญ: กรุณาตรวจสอบการทำธุรกรรมและข้อมูลการชำระเงินระหว่างการทำธุรกรรมสกุลเงิน GIAC โปรดระมัดระวังเกี่ยวกับข้อมูล",
		Ko_KR: "중요한 알림: GIAC 화폐 거래 시 거래정보와 결제정보를 확인하시기 바랍니다. 해당 정보에 주의를 기울이기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTE4,
		Zh_CN: "本邮件由GIAC系统自动发出，请勿直接回复。",
		En_US: "This email is sent automatically by GIAC system. Please do not reply directly.",
		Zh_HK: "本郵件由GIAC系統自動發出，請勿直接回復。",
		Ja_JP: "本メールはGIACシステムが自動的に送信しちえます。直接返信はしないでください。",
		Vi_VN: "Email này được tự động gửi từ hệ thống của GIAC. Vui lòng không trả lời trực tiếp.",
		Ru_RU: "Это сообщение отправлено автоматически системой GIAC. На него не нужно отвечать.",
		Pt_BR: "Este e-mail é enviado automaticamente pelo sistema GIAC. Não responda diretamente.",
		Th_TH: "อีเมลฉบับนี้ส่งมาจากระบบ GIAC โดยอัตโนมัติ",
		Ko_KR: "본 이메일은 GIAC 시스템에서 자동으로 발송되는",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_NOTE5,
		Zh_CN: "如果您有任何疑问或建议，请联系我们，E-mail:{{.SupportEmail}}",
		En_US: "If you have any questions or suggestions, please contact us, E-mail:{{.SupportEmail}}",
		Zh_HK: "如果您有任何疑問或建議，請聯繫我們，E-mail:{{.SupportEmail}}",
		Ja_JP: "ご質問やご意見がある場合はご連絡ください E-mail:{{.SupportEmail}}",
		Vi_VN: "Nếu bạn có bất kỳ câu hỏi hoặc đề xuất nào, vui lòng liên hệ với chúng tôi. E-mail:{{.SupportEmail}}",
		Ru_RU: "Если у вас имеются вопросы или предложения, свяжитесь с нами. E-mail:{{.SupportEmail}}",
		Pt_BR: "Se tiver qualquer dúvida ou sugestão, entre em contato conosco. E-mail:{{.SupportEmail}}",
		Th_TH: "หากคุณมีข้อสงสัยหรือคำแนะนำ กรุณาติดต่อเรา E-mail:{{.SupportEmail}}",
		Ko_KR: "궁금하신 사항이나 제안하고 싶으신 사항이 있으시면 당사에 문의해 주십시오. E-mail:{{.SupportEmail}}",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    EMAIL_TIPS,
		Zh_CN: "请登录您的 APP账户，及时查看最新消息！",
		En_US: "Please login to your GIAC APP account and check the latest news in time!",
		Zh_HK: "請登入您的GIAC APP帳戶，及時查看最新消息！",
		Ja_JP: "ご使用のGIAC アプリのアカウントを登録してください。いつでも最新情報がチェックできます。",
		Vi_VN: "Vui lòng đăng nhập vào tài khoản GIAC APP của bạn để nhận thông tin mới nhất.",
		Ru_RU: "Войдите в свой лицевой счет в приложении GIAC для получения последних новостей.",
		Pt_BR: "Fazer o login na sua conta GIAC APP para verificar as últimas notícias.",
		Th_TH: "กรุณาล็อกอินเข้าสู่บัญชีแอปฯ GIAC ของคุณเพื่อดูข่าวสารล่าสุด",
		Ko_KR: "가장 최근 뉴스를 확인하려면 GIAC 어플 계정으로 로그인하시기 바랍니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_1,
		Zh_CN: "银行卡",
		En_US: "Bank card",
		Zh_HK: "銀行卡",
		Ja_JP: "銀行カード",
		Vi_VN: "Thẻ ngân hàng",
		Ru_RU: "Банковская карта",
		Pt_BR: "Cartão do banco",
		Th_TH: "บัตรธนาคาร",
		Ko_KR: "은행 카드",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_2,
		Zh_CN: "微信",
		En_US: "WeChat",
		Zh_HK: "WeChat",
		Ja_JP: "WeChat",
		Vi_VN: "WeChat",
		Ru_RU: "WeChat",
		Pt_BR: "WeChat",
		Th_TH: "WeChat",
		Ko_KR: "WeChat",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_3,
		Zh_CN: "支付宝",
		En_US: "Alipay",
		Zh_HK: "支付寶",
		Ja_JP: "Alipay",
		Vi_VN: "Alipay",
		Ru_RU: "Alipay",
		Pt_BR: "Alipay",
		Th_TH: "Alipay",
		Ko_KR: "Alipay",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_4,
		Zh_CN: "LINEpay",
		En_US: "LINEpay",
		Zh_HK: "LINEpay",
		Ja_JP: "LINEpay",
		Vi_VN: "LINEpay",
		Ru_RU: "LINEpay",
		Pt_BR: "LINEpay",
		Th_TH: "LINEpay",
		Ko_KR: "LINEpay",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_5,
		Zh_CN: "PayPal",
		En_US: "PayPal",
		Zh_HK: "PayPal",
		Ja_JP: "PayPal",
		Vi_VN: "PayPal",
		Ru_RU: "PayPal",
		Pt_BR: "PayPal",
		Th_TH: "PayPal",
		Ko_KR: "PayPal",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ITEM_USE_BANK_ERROR,
		Zh_CN: "有广告正在使用此收款方式，无法删除",
		En_US: "There are advertisements that are using this method of collection and cannot be deleted.",
		Zh_HK: "有廣告正在使用此收款管道，無法删除",
		Ja_JP: "広告がこの料金受領方法を使用しているので、削除できません。",
		Vi_VN: "Một số quảng cáo đang sử dụng phương thức thanh toán này và không thể bị xóa.",
		Ru_RU: "Рекламные объявления, использующие данный способ оплаты, невозможно удалить",
		Pt_BR: "Há anúncios que usam esse método de pagamento e não podem ser excluídos",
		Th_TH: "มีโฆษณาต่าง ๆ ที่ใช้วิธีการชำระเงินนี้และไม่สามารถลบทิ้งได้",
		Ko_KR: "해당 결제 방법을 사용하고 있는 광고가 있어 본 결제 방법을 제거할 수 없습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    BANK_METHOD_LIMIT,
		Zh_CN: "该收款方式只能添加一个账户",
		En_US: "Only one account can be added to the collection method.",
		Zh_HK: "該收款管道只能添加一個帳戶",
		Ja_JP: "この受領方法では、1つの口座のみ追加できます",
		Vi_VN: "Phương thức thanh toán này chỉ có thể thêm một tài khoản.",
		Ru_RU: "Для этого способа оплаты можно добавить только один лицевой счет.",
		Pt_BR: "Este método de pagamento só pode adicionar uma conta.",
		Th_TH: "การชำระเงินวิธีนี้สามารถเพิ่มได้เพียงหนึ่งบัญชีเท่านั้น",
		Ko_KR: "해당 결제 방법은 한 계정에만 추가될 수 있습니다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    CANCLE_TOO_MUCH,
		Zh_CN: "今日撤单过多，无法继续交易",
		En_US: "Too many withdrawals today to continue trading",
		Zh_HK: "今日撤單過多，無法繼續交易",
		Ja_JP: "今日はキャンセルが多すぎて、取引を続けられません。",
		Vi_VN: "Hôm nay quá nhiều rút tiền để tiếp tục giao dịch",
		Ru_RU: "сегодня отменить слишком много заказов, невозможно продолжить сделку",
		Pt_BR: "Muitas retiradas hoje para continuar a negociação",
		Th_TH: "วันนี้มีการยกเลิกคำสั่งมากเกินไปที่จะดำเนินการต่อ",
		Ko_KR: "오늘 취소 신청이 너무 많아서, 계속 거래할 수 없다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    AMOUNT_ERROR,
		Zh_CN: "数额不在区间内，{{.min}}~{{.max}}",
		En_US: "The amount is not in the range，{{.min}}~{{.max}}",
		Zh_HK: "數額不在區間內，{{.min}}~{{.max}}",
		Ja_JP: "金額は区間内ではないです，{{.min}}~{{.max}}",
		Vi_VN: "Số tiền đó không nằm trong phạm vi，{{.min}}~{{.max}}",
		Ru_RU: "сумма не находится в интервале，{{.min}}~{{.max}}",
		Pt_BR: "A quantidade não está no intervalo，{{.min}}~{{.max}}",
		Th_TH: "วจำนวนไม่ได้อยู่ในช่วง，{{.min}}~{{.max}}",
		Ko_KR: "액수가 구간 내에 없다，{{.min}}~{{.max}}",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    AMOUNT_TOO_MUCH,
		Zh_CN: "超出限额",
		En_US: "Exceeding quota",
		Zh_HK: "超出限額",
		Ja_JP: "限度を超える",
		Vi_VN: "Vượt quá ngạch",
		Ru_RU: "превысить лимит",
		Pt_BR: "Quota excedentária",
		Th_TH: "เกินโควต้า",
		Ko_KR: "한도를 초과하다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_PRICE_LESS,
		Zh_CN: "价格过低",
		En_US: "price less than min",
		Zh_HK: "價格過低",
		Ja_JP: "価格が低すぎる",
		Vi_VN: "Giá thì quá thấp.",
		Ru_RU: "низкая цена",
		Pt_BR: "O preço é Muito baixo.",
		Th_TH: "ราคาต่ำเกินไป",
		Ko_KR: "가격이 너무 낮다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_PRICE_GREATER,
		Zh_CN: "价格太高",
		En_US: "price greater than max",
		Zh_HK: "價格太高",
		Ja_JP: "値段が高すぎる",
		Vi_VN: "Giá quá đắt rồi.",
		Ru_RU: "цена слишком высокая",
		Pt_BR: "O preço é Muito alto.",
		Th_TH: "ราคาสูงเกินไป",
		Ko_KR: "가격이 너무 비싸다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_PRICE_NOT_LAWFUL,
		Zh_CN: "价格不合法",
		En_US: "price is not lawful",
		Zh_HK: "價格不合法",
		Ja_JP: "価格は非合法です",
		Vi_VN: "Giá bất hợp pháp",
		Ru_RU: "неправильные цены",
		Pt_BR: "Preço ilegal",
		Th_TH: "ราคาผิดกฎหมาย",
		Ko_KR: "가격은 비합법적이다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_PRICE_FLUCTUATE_BIG,
		Zh_CN: "价格波动太大",
		En_US: "prices fluctuate too much",
		Zh_HK: "價格波動太大",
		Ja_JP: "価格の変動が大きすぎる",
		Vi_VN: "Quá nhiều mức bất ổn",
		Ru_RU: "колебания цен слишком велики",
		Pt_BR: "Demasiada volatilidade DOS preços",
		Th_TH: "ความผันผวนของราคามากเกินไป",
		Ko_KR: "가격 파동 이 너무 크다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_QUANTITY_LESS,
		Zh_CN: "数量太小",
		En_US: "quantity less than min",
		Zh_HK: "數量太小",
		Ja_JP: "数量が小さすぎる",
		Vi_VN: "Số lượng đó quá nhỏ.",
		Ru_RU: "количество слишком мало",
		Pt_BR: "A quantidade é Muito pequena.",
		Th_TH: "จำนวนน้อยเกินไป",
		Ko_KR: "수량이 너무 작다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_QUANTITY_GREATER,
		Zh_CN: "数量太多",
		En_US: "quantity greater than max",
		Zh_HK: "數量太多",
		Ja_JP: "数量が多すぎる",
		Vi_VN: "Quá nhiều lượng",
		Ru_RU: "слишком много",
		Pt_BR: "Quantidades demais.",
		Th_TH: "จำนวนที่มากเกินไป",
		Ko_KR: "수량이 너무 많다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_QUANTITY_NOT_LAWFUL,
		Zh_CN: "数量不合法",
		En_US: "quantity is not lawful",
		Zh_HK: "數量不合法",
		Ja_JP: "数量が適法でない",
		Vi_VN: "Số lượng bất hợp pháp",
		Ru_RU: "неправильное количество",
		Pt_BR: "Quantidade ilegal",
		Th_TH: "ปริมาณที่ผิดกฎหมาย",
		Ko_KR: "수량이 비합법적이다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_QUANTITY_NOT_LAWFUL,
		Zh_CN: "数量不合法",
		En_US: "quantity is not lawful",
		Zh_HK: "數量不合法",
		Ja_JP: "数量が適法でない",
		Vi_VN: "Số lượng bất hợp pháp",
		Ru_RU: "неправильное количество",
		Pt_BR: "Quantidade ilegal",
		Th_TH: "ปริมาณที่ผิดกฎหมาย",
		Ko_KR: "수량이 비합법적이다.",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    ORDER_FILTER_QUOTEASSET_NOT_LAWFUL,
		Zh_CN: "行情币数量不合法",
		En_US: "quote asset is not lawful",
		Zh_HK: "行情幣數量不合法",
		Ja_JP: "貨幣の数は合法ではないです。",
		Vi_VN: "Số tiền tệ bất hợp pháp",
		Ru_RU: "Незаконный оборот валюты",
		Pt_BR: "Montante ilegal de moeda de Mercado",
		Th_TH: "ปริมาณของสกุลเงินที่ไม่ถูกต้อง",
		Ko_KR: "시세화의 수량이 합법적이지 않다",
	})

	i18ns.AddI18n(utils.I18n{
		Id:    SYMBOL_ERROR,
		Zh_CN: "无此交易对",
		En_US: "No such transaction pair",
		Zh_HK: "無此交易對",
		Ja_JP: "この取引ペアはありません",
		Vi_VN: "Không có cặp giao dịch đó",
		Ru_RU: "Нет такой сделки",
		Pt_BR: "Nenhum par de transações",
		Th_TH: "ไม่มีคู่ของธุรกรรมนี้",
		Ko_KR: "이 교역이 없다",
	})

	if err := i18nfile.WriteToFile(i18ns.Items, utils.GetAppPath()+"/i18n"); err != nil {
		panic(err.Error())
	}
}
