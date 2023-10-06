import { I18n } from "@/lib/i18n";

const ja: I18n = {
  app: "LGTM Generator",
  description:
    "シンプルな LGTM 画像生成サービスです。 LGTM とは「Looks Good To Me」の略で、コードレビューのときなどに用いられるネットスラングの一種です。",

  lgtm: "LGTM",
  searchImage: "画像検索",
  favorite: "お気に入り",

  loadMore: "もっと見る",
  upload: "アップロード",
  confirmGeneration: "この画像で LGTM を生成しますか？",
  generate: "生成",
  send: "送信",
  supplement: "( 任意 ) 補足",
  illegal: "法律違反 ( 著作権侵害、プライバシー侵害、名誉毀損等 )",
  inappropriate: "不適切なコンテンツ",
  other: "その他",
  keyword: "キーワード",
  cancel: "キャンセル",
  random: "ランダムに表示",
  reload: "再読み込み",

  privacyPolicy: "プライバシーポリシー",
  usagePrecautions: "利用上の注意",
  pleaseReadUsagePrecautions: (locale: string) => (
    <>
      LGTM を生成する前に
      <a
        className="text-primary-main underline"
        href={`/${locale}/usage-precautions`}
        target="_blank"
        rel="noopener noreferrer"
      >
        利用上の注意
      </a>
      をお読みください。
    </>
  ),

  noFavorites: "お気に入りした LGTM はありません",
  successToGenerate: "LGTM を生成しました",
  failedToGenerate: "LGTM の生成に失敗しました",
  successToSend: "送信しました",
  failedToSend: "送信に失敗しました",

  fileTooLarge: "ファイルサイズが大きすぎます",
  unsupportedImageFormat: "サポートされていない画像形式です",
  copiedToClipboard: "クリップボードにコピーしました",
  rateLimitReached:
    "リクエストの上限に達しました。\nしばらく待ってから再度お試しください。",

  aboutTrafficAnalysisTool: "アクセス解析ツールについて",
  aboutTrafficAnalysisToolContent: (
    <>
      当サイトでは、 Google によるアクセス解析ツール「Google
      アナリティクス」を利用しています。この Google
      アナリティクスはトラフィックデータの収集のために Cookie
      を使用しています。このトラフィックデータは匿名で収集されており、個人を特定するものではありません。この機能は
      Cookie
      を無効にすることで収集を拒否することが出来ますので、お使いのブラウザの設定をご確認ください。この規約に関して、詳しくは{" "}
      <a
        className="text-primary-main underline"
        href="https://marketingplatform.google.com/about/analytics/terms/jp/"
        target="_blank"
        rel="noopener noreferrer"
      >
        {" "}
        Google アナリティクス利用規約
      </a>
      を参照してください。
    </>
  ),

  aboutPrivacyPolicyChange: "プライバシーポリシー変更について",
  aboutPrivacyPolicyChangeContent:
    "当サイトは、個人情報に関して適用される日本の法令を遵守するとともに、本ポリシーの内容を適宜見直しその改善に努めます。修正された最新のプライバシーポリシーは常に本ページにて開示されます。",

  usagePrecautionsItems: [
    "サービスを利用して生成された画像に関する一切の責任はご利用者様に負担いただくものとします。ご利用者様が生成した画像に関し、第三者が損害を被った場合、運営者はご利用者様に代わっての責任は一切負いません。",
    "本サービスを利用して生成された画像はインターネット上に公開されます。元画像の著作権や関連法規に注意してください。公序良俗に反する画像や違法な画像を作成しないでください。これらの画像、その他運営者が不適切と判断した画像は予告無しに削除することがあります。",
    "過剰な数のリクエストを送信してサービスに負荷をかける行為はおやめください。",
    "その他、悪質な利用方法が確認された場合、特定のご利用者様を予告無しにアクセス禁止にすることがあります。",
  ],
} as const;

export default ja;
