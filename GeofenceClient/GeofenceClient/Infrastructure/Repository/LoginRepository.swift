import Foundation

public protocol ILoginRepository {
    func login(id: String, pass: String)
    func logout(id: String)
}

class LoginRepository: ILoginRepository {
    func login(id: String, pass: String) {
        
    }
    
    func logout(id: String) {
        
    }
}
