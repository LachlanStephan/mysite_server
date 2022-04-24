<?php
    function insertNewSection($data) 
    {
        global $conn;

        $sql = "INSERT INTO section (title, description)
        VALUES
        (:t, :d)";

        $stmt = $conn->prepare($sql);
        $stmt->execute([
            ':t' => $data['title'],
            ':d' => $data['description'],
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

        $sql = "SELECT * FROM section ORDER BY section_id DESC"; // add limit later 10?

        $stmt = $conn->prepare($sql);
        $stmt->execute();
        
        $rows = $stmt->fetchAll();

        return $rows;
    }

    function updateSection($data) 
    {
        global $conn;

        $sql = "UPDATE section SET title = :t, description = :d WHERE section_id = :id";

        $stmt = $conn->prepare($sql);
        $stmt->execute([
            ':t' => $data['title'],
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