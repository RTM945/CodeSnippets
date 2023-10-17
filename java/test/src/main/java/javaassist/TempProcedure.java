package javaassist;

public class TempProcedure {

    private Msg msg;
	
	public TempProcedure(Msg msg) {
		this.msg = msg;
	}
	
	public TempProcedure() {
	}
	
	protected boolean process() throws Exception {
		return true;
	}

	public void queueMsg(final Msg msg) throws Exception {
		System.out.println("queueMsg " + msg.getType());
		process();
	}
}
