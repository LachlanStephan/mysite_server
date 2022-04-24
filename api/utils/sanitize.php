<?php
    function sanitizeData($d)
    {
            $d = trim($d);
            $d = stripslashes($d);
            $d = htmlspecialchars($d);
            $d = htmlentities($d);
            return $d;
    }
?>