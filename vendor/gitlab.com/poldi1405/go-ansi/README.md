# go-ansi

Write pretty terminal output the easy way.

Want your text green? `ansi.Green("I am Shrek")`

Want it black on a red ground? `ansi.RedBG(ansi.Black("You can just combine them"))`

Moving around is as simple as Specifying direction `ansi.Up()`

## ANSI Support

| Terminal   | Colors             | BG Colors          | Styles               | Movements          | Saving Positions   |
|------------|--------------------|--------------------|----------------------|--------------------|--------------------|
| Kitty      | :heavy_check_mark: | :heavy_check_mark: | :heavy_minus_sign:¹² | :heavy_check_mark: | :heavy_check_mark: |
| Alacritty  | :heavy_check_mark: | :heavy_check_mark: | :heavy_minus_sign:¹  | :heavy_check_mark: | :heavy_check_mark: |
| xTerm      | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:³  | :heavy_check_mark: | :heavy_check_mark: |
| st         | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:³  | :heavy_check_mark: | :heavy_check_mark: |
| guake      | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:³  | :heavy_check_mark: | :heavy_check_mark: |
| urxvt      | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:²⁴ | :heavy_check_mark: | :heavy_check_mark: |
| terminator | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:²³ | :heavy_check_mark: | :heavy_check_mark: |
| termite    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:²³ | :heavy_check_mark: | :heavy_check_mark: |

¹) Blinking not supported
²) Conceal not supported
³) BUG: fast blinking 
⁴) Strikethrough not working

*this list may be inaccurate as it was created after an exhausting day of work at a time where I wasn't up to much*
