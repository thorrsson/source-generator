using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Authentication;

namespace GitHubAuthentication;

// TODO: Consider implementing `Bearer` authentication scheme as a separate class
// TODO: Consider implementing `Basic` authentication scheme as a separate class
// TODO: Consider implementing `Anonymous` authentication scheme as a separate class
// TODO: Consider implementing `AccessToken/Device` authentication scheme as a separate class
// TODO: Consider implementing `GitHub Apps` authentication scheme as a separate class
public class GitHubTokenAuthenticationProvider : IAuthenticationProvider
{
  private const string AuthorizationHeaderKey = "Authorization";

  public IAuthenticationProvider TokenAuthProvider {get; private set;}
  public string ClientId { get; set; }
  public string Token { get; set; }

  /// <summary>
  /// 
  /// </summary>
  /// <param name="clientId"></param>
  /// <param name="token"></param>
  /// <exception cref="ArgumentNullException"></exception>
  public GitHubTokenAuthenticationProvider(string clientId, string token)
	{
    if (string.IsNullOrEmpty(clientId)) throw new ArgumentNullException(nameof(clientId));
    if (string.IsNullOrEmpty(token)) throw new ArgumentNullException(nameof(token));

    ClientId = clientId;
    Token = token;

    TokenAuthProvider = this;
  }

  /// <summary>
  /// Authenticates the application request.
  /// This currently acts as a synchronous (no await) method - but will change as we do more work on the authentication library.
  /// </summary>
  /// <param name="request"></param>
  /// <param name="additionalAuthenticationContext"></param>
  /// <param name="cancellationToken"></param>
  /// <returns></returns>
  /// <exception cref="ArgumentNullException"></exception>
  public async Task AuthenticateRequestAsync(RequestInformation request, Dictionary<string, object>? additionalAuthenticationContext = null, CancellationToken cancellationToken = default)
  {

    if(request == null) throw new ArgumentNullException(nameof(request));
    
    // This needs to be abstracted out into a configuration option but still set on each request
    request.Headers.Add("X-GitHub-Api-Version", "2022-11-28");
    request.Headers.Add("User-Agent", ClientId);

    if (Token == string.Empty)
    {
      throw new ArgumentNullException(nameof(Token));
    }

    if(!request.Headers.ContainsKey(AuthorizationHeaderKey))
    {
      request.Headers.Add(AuthorizationHeaderKey, Token);
    }
  }
}