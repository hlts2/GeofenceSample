import Foundation

protocol HTTPRequest {
    var baseURL: String { get }
    var path   : String { get }
    var method : HTTPMethod { get set }
    var params : [String: String] { get set}
}

extension HTTPRequest {
    
    func createRequest() {
        
    }
}
