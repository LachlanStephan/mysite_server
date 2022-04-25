<?php
    function emailContactMessage($data) 
    {
        $recipient = "ljstephan116@gmail.com";
        $subject = $data['subject'];
        $msg = $data['message'];
        $sent = mail($recipient, $subject, $msg);

        if ($sent) {
            return true;
        }
        if (!$sent) {
            return false;
        }
    }

?>