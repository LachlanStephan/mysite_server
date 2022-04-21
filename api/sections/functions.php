<?php
    function insertNewSection($data) 
    {
        global $conn;

        $sql = "INSERT INTO section (title, content)
        VALUES
        (:t, :c)";

        $stmt = $conn->prepare($sql);
        $stmt->execute([
            ':t' => $data['title'],
            ':c' => $data['content'],
        ]);

        if (!empty($conn->lastInsertId())) {
            return true;
        } else {
            return false;
        }
    }

    function getSections() 
    {
        global $conn;

        $sql = "SELECT * FROM section ORDER BY section_id DESC"; // add limit later

        $stmt = $conn->prepare($sql);
        $stmt->execute();
        
        $rows = $stmt->fetchAll();

        return $rows;
    }

?>