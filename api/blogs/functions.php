<?php
    function insertNewBlog($data)
    {
        global $conn;

        $sql = "INSERT INTO blog (title, description, content)
        VALUES
        (:t, :d, :c) ";
        
        $stmt = $conn->prepare($sql);
        $stmt->execute([
            ':t' => $data['title'],
            ':d' => $data['description'],
            ':c' => $data['content'],
        ]);

        if (!empty($conn->lastInsertId())) {
            return true;
        } else {
            return false;
        }
    }

    function getBlogs()
    {
        global $conn;

        $sql = "SELECT * FROM blog ORDER BY blog_id DESC"; // will need to limit at some point

        $stmt = $conn->prepare($sql);
        $stmt->execute();

        $rows = $stmt->fetchAll();

        return $rows;
    }

    function updateBlog($data) 
    {
        global $conn;

        $sql = "UPDATE blog SET title = :t, content = :c, description = :d WHERE blog_id = :id";

        $stmt = $conn->prepare($sql);
        $stmt->execute([
            ':t' => $data['title'],
            ':c' => $data['content'],
            ':d' => $data['description'],
            ':id' => $data['id'],
        ]);

        if ($stmt->rowCount() > 0) {
            return true;
        } else {
            return false;
        }
    }

?>