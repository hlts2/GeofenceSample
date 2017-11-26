import UIKit

protocol ILoginVC {
    
}

class LoginViewController: UIViewController {
    
    lazy var presenter = {
        return LoginVCPresenter(viewInput: self)
    }()
    
    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}

extension LoginViewController: ILoginVC {
    
}
