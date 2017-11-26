import Foundation

public protocol ILoginUseCase {
    func login(id: String, pass: String)
    func logout(id: String)
}

class LoginUseCase: ILoginUseCase {
    
    let repository: ILoginRepository = {
        return LoginRepository()
    }()
    
    public func login(id: String, pass: String) {
        repository.login(id: id, pass: pass)
    }
    
    public func logout(id: String) {
        repository.logout(id: id)
    }
}
