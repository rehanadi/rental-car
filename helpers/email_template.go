package helpers

const WelcomeEmailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333333;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        .header {
            background-color: #4A90E2;
            color: white;
            padding: 20px;
            text-align: center;
            border-radius: 5px;
        }
        .content {
            padding: 20px;
            background-color: #ffffff;
        }
        .footer {
            text-align: center;
            padding: 20px;
            font-size: 0.9em;
            color: #666666;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            background-color: #4A90E2;
            color: white;
            text-decoration: none;
            border-radius: 5px;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>Welcome to Our Rental Car Service!</h1>
    </div>
    
    <div class="content">
        <p>Dear %s,</p>
        
        <p>Welcome to our Rental Car Service! We're excited to have you on board.</p>
        
        <p>Your account has been successfully created. You can now log in and start exploring our services.</p>
        
        <p>If you have any questions or need assistance, our support team is here to help!</p>
        
        <p>Best regards,<br>
        The Rental Car Team</p>
    </div>
    
    <div class="footer">
        <p>This email was sent to %s</p>
        <p>© 2024 Rental Car Service. All rights reserved.</p>
    </div>
</body>
</html>
`

const WelcomeEmailPlainTemplate = `
Dear %s,

Welcome to our Rental Car Service! We're excited to have you on board.

Your account has been successfully created. You can now log in and start exploring our services.

If you have any questions, please don't hesitate to contact our support team.

Best regards,
The Rental Car Team
`
