import UIKit

protocol ITopVC {
    
}

class TopViewController: UIViewController {
    
    lazy var presenter = {
        return TopVCPresenter(viewInput: self)
    }()
    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
}

extension TopViewController: ITopVC {
    
}
