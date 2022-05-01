<?php
    require("../../index.php");
    require("./functions.php");

    $id = $_GET['id'];
    $theme = $_GET['theme'];
    $blogContent = getRequestedBlog($id);

?> 

<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
        <link rel="stylesheet" href="./style.css">
    </head>
    <body 
       class = "<?php
            if ($theme === 'lightMode') { echo 'lightMode'; };
            if ($theme === 'dark') { echo ''; };
        ?>"
    >
        <main>
            <br>
            <a href="https://ljstephan.dev" class="blog_back_btn"><- Back</a>
            <?php
                echo $blogContent['content'];
            ?>
        </main>
    </body>
    </html>