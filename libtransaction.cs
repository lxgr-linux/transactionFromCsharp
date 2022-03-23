using System.Runtime.InteropServices;

namespace libTransaction {

  public class Transaction {
    [DllImport("transaction.so", EntryPoint="makeConfirmMatchRequest")]
    public static extern void msgConfirmMatch(string creator, int matchId, string rawOutcome);
  }
}
